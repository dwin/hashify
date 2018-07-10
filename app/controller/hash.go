package controller

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"net/http"
	"strings"

	"golang.org/x/crypto/blake2s"

	"github.com/labstack/echo"
	blake2bminio "github.com/minio/blake2b-simd"
	"github.com/minio/sha256-simd"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
)

func ComputeHash(c echo.Context) error {
	var h hash.Hash
	algorithm := strings.ToUpper(c.Param("algo"))

	// Determine Hash Method
	switch algorithm {
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
		return c.String(http.StatusNotFound, "Invalid Path")
	}

	// Check Request Method
	if c.Request().Method == http.MethodGet {
		val := c.QueryParam("value")
		_, err := h.Write([]byte(val))
		if err != nil {
			return err
		}
	}
	if c.Request().Method == http.MethodPost {
		io.Copy(h, c.Request().Body)
	}
	j, err := json.Marshal(HashResp{
		Digest: hex.EncodeToString(h.Sum(nil)),
		Type:   algorithm,
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
