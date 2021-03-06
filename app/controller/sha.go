package controller

import (
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/minio/sha256-simd"

	"github.com/labstack/echo"
)

func HashSHA1(c echo.Context) error {
	hash := sha1.New()
	// Check Request Method
	if c.Request().Method == http.MethodGet {
		val := c.QueryParam("value")
		_, err := hash.Write([]byte(val))
		if err != nil {
			return err
		}
	}
	if c.Request().Method == http.MethodPost {
		io.Copy(hash, c.Request().Body)
	}
	j, err := json.Marshal(HashResp{
		Digest: hex.EncodeToString(hash.Sum(nil)),
		Type:   "SHA1",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func HashSHA256(c echo.Context) error {
	hash := sha256.New()
	// Check Request Method
	if c.Request().Method == http.MethodGet {
		val := c.QueryParam("value")
		_, err := hash.Write([]byte(val))
		if err != nil {
			return err
		}
	}
	if c.Request().Method == http.MethodPost {
		io.Copy(hash, c.Request().Body)
	}
	j, err := json.Marshal(HashResp{
		Digest: hex.EncodeToString(hash.Sum(nil)),
		Type:   "SHA-256",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func HashSHA384(c echo.Context) error {
	hash := sha512.New384()
	// Check Request Method
	if c.Request().Method == http.MethodGet {
		val := c.QueryParam("value")
		_, err := hash.Write([]byte(val))
		if err != nil {
			return err
		}
	}
	if c.Request().Method == http.MethodPost {
		io.Copy(hash, c.Request().Body)
	}
	j, err := json.Marshal(HashResp{
		Digest: hex.EncodeToString(hash.Sum(nil)),
		Type:   "SHA-384",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func HashSHA512(c echo.Context) error {
	hash := sha512.New()
	// Check Request Method
	if c.Request().Method == http.MethodGet {
		val := c.QueryParam("value")
		_, err := hash.Write([]byte(val))
		if err != nil {
			return err
		}
	}
	if c.Request().Method == http.MethodPost {
		io.Copy(hash, c.Request().Body)
	}
	j, err := json.Marshal(HashResp{
		Digest: hex.EncodeToString(hash.Sum(nil)),
		Type:   "SHA-512",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
