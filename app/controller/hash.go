package controller

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/blake2s"

	"github.com/labstack/echo"
	blake2bminio "github.com/minio/blake2b-simd"
	"github.com/minio/highwayhash"
	"github.com/minio/sha256-simd"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
)

type errorMsg struct {
	Error string
}

func ComputeHash(c echo.Context) error {
	var h hash.Hash
	algorithm := strings.ToUpper(c.Param("algo"))
	format := strings.ToLower(c.Param("format"))
	var keyHex string

	if len(algorithm) < 1 {
		return c.JSONP(http.StatusNotFound, "", &BasicError{
			Error: "Path does not match available endpoint (algorithm), see API documentation",
		})
	}
	if len(format) < 1 {
		return c.JSONP(http.StatusNotFound, "", &BasicError{
			Error: "Path does not match available endpoint (output), see API documentation",
		})
	}
	// Determine Hash Method
	switch algorithm {
	case "HIGHWAY":
		key, err := parseHighwayHashKey(c)
		if err != nil {
			e := BasicError{
				Error: "Invalid Hex Value for parameter \"key\", must provide key as query param or header \"X-Hashify-Key\"",
			}
			return c.JSON(http.StatusBadRequest, e)
		}
		// Check Key Length
		if len(key) != 32 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  "HighwayHash key parameter must be 32 bytes",
				"length": fmt.Sprintf("%v", len(key)),
			})
		}
		hash, err := highwayhash.New(key)
		if err != nil {
			return err
		}
		h = hash
		algorithm = "HighwayHash-256"
		keyHex = hex.EncodeToString(key)
	case "HIGHWAY-64":
		key, err := parseHighwayHashKey(c)
		if err != nil {
			e := BasicError{
				Error: "Invalid Hex Value for parameter \"key\", must provide key as query param or header \"X-Hashify-Key\"",
			}
			return c.JSON(http.StatusBadRequest, e)
		}
		// Check Key Length
		if len(key) != 32 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  "HighwayHash key parameter must be 32 bytes",
				"length": fmt.Sprintf("%v", len(key)),
			})
		}
		hash, err := highwayhash.New(key)
		if err != nil {
			return err
		}
		h = hash
		algorithm = "HighwayHash-64"
		keyHex = hex.EncodeToString(key)
	case "HIGHWAY-128":
		key, err := parseHighwayHashKey(c)
		if err != nil {
			e := BasicError{
				Error: "Invalid Hex Value for parameter \"key\", must provide key as query param or header \"X-Hashify-Key\"",
			}
			return c.JSON(http.StatusBadRequest, e)
		}
		// Check Key Length
		if len(key) != 32 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  "HighwayHash key parameter must be 32 bytes",
				"length": fmt.Sprintf("%v", len(key)),
			})
		}
		hash, err := highwayhash.New(key)
		if err != nil {
			return err
		}
		h = hash
		algorithm = "HighwayHash-128"
		keyHex = hex.EncodeToString(key)
	case "MD4":
		h = md4.New()
	case "MD5":
		h = md5.New()
	case "SHA1":
		h = sha1.New()
	case "SHA256":
		h = sha256.New()
	case "SHA384":
		h = sha512.New384()
	case "SHA512":
		h = sha512.New()
	case "SHA512-256":
		h = sha512.New512_256()
	case "SHA3-256":
		h = sha3.New256()
	case "SHA3-384":
		h = sha3.New384()
	case "SHA3-512":
		h = sha3.New512()
	case "BLAKE2B-256":
		h = blake2bminio.New256()
	case "BLAKE2B-384":
		hash, err := blake2b.New384(nil)
		if err != nil {
			return err
		}
		h = hash
	case "BLAKE2B-512":
		h = blake2bminio.New512()
	case "BLAKE2S-128":
		hash, err := blake2s.New128(nil)
		if err != nil {
			return err
		}
		h = hash
	case "BLAKE2S-256":
		hash, err := blake2s.New256(nil)
		if err != nil {
			return err
		}
		h = hash
	default:
		return c.JSONP(http.StatusNotFound, "", &BasicError{
			Error: "Path does not match available endpoint, see API documentation",
		})
	}

	// Check Request Method
	if c.Request().Method == http.MethodGet {
		val := c.QueryParam("value")
		_, err := h.Write([]byte(val))
		if err != nil {
			return err
		}
	}
	// Handle Form file
	if c.Request().Header.Get("Content-Type") == "multipart/form-data" || c.Request().Header.Get("X-Hashify-Process") == "multipart/form-data" {
		file, err := c.FormFile("file")
		if err != nil {
			log.Printf("get form file error: %s\n", err)
			eMsg := errorMsg{
				Error: "File object must be named 'file' when using client headers of ('Content-Type' || 'X-Hashify-Process' = 'multipart/form-data')",
			}
			return c.JSONPretty(http.StatusBadRequest, eMsg, " ")
		}

		src, err := file.Open()
		if err != nil {
			log.Printf("open form file error: %s\n", err)
			return err
		}
		defer src.Close()
		io.Copy(h, src)
	}

	if c.Request().Method == http.MethodPost && c.Request().Header.Get("Content-Type") != "multipart/form-data" {
		io.Copy(h, c.Request().Body)
	}
	// Get Digest in requested format
	var digest string
	switch format {
	case "base32":
		digest = base32.StdEncoding.EncodeToString(h.Sum(nil))
	case "base64":
		digest = base64.StdEncoding.EncodeToString(h.Sum(nil))
	case "base64url":
		digest = base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	case "hex":
		digest = hex.EncodeToString(h.Sum(nil))
	default:
		return c.JSONP(http.StatusNotFound, "", &BasicError{
			Error: "Path does not match available endpoint (output), see API documentation",
		})
	}

	j, err := json.Marshal(HashResp{
		Digest:    digest,
		DigestEnc: format,
		Type:      algorithm,
		Key:       keyHex,
	})
	if err != nil {
		log.Printf("json marshal hash response error: %s\n", err)
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func hashString(h hash.Hash, plaintext string, algorithm string) (resp []byte, err error) {

	return json.Marshal(HashResp{
		Digest: hex.EncodeToString(h.Sum([]byte(plaintext))),
		Type:   algorithm,
		Key:    "",
	})
}
