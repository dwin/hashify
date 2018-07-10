package controller

import (
	"encoding/hex"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/sha3"
)

func HashSHA3_256(c echo.Context) error {
	hash := sha3.New256()
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
		Type:   "SHA3-256",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func HashSHA3_384(c echo.Context) error {
	hash := sha3.New384()
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
		Type:   "SHA3-384",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func HashSHA3_512(c echo.Context) error {
	hash := sha3.New512()
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
		Type:   "SHA3-512",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
