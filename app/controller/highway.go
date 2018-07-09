package controller

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/minio/highwayhash"
)

type BasicError struct {
	Error string
}

func HashHighwayHash(c echo.Context) error {
	var key []byte
	if c.QueryParam("key") == "random" {
		k, err := randKey(32)
		if err != nil {
			return err
		}
		key = k
	} else {
		k, err := hex.DecodeString(c.QueryParam("key"))
		if err != nil {
			e := BasicError{
				Error: "Invalid Hex Value for parameter \"key\"",
			}
			return c.JSON(http.StatusBadRequest, e)
		}
		key = k
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
		Type:   "HighwayHash-256",
		Key:    hex.EncodeToString(key),
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func HashHighwayHash64(c echo.Context) error {
	var key []byte
	if c.QueryParam("key") == "random" {
		k, err := randKey(32)
		if err != nil {
			return err
		}
		key = k
	} else {
		k, err := hex.DecodeString(c.QueryParam("key"))
		if err != nil {
			e := BasicError{
				Error: "Invalid Hex Value for parameter \"key\"",
			}
			return c.JSON(http.StatusBadRequest, e)
		}
		key = k
	}
	// Check Key Length
	if len(key) != 32 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "HighwayHash key parameter must be 32 bytes",
		})
	}
	hash, err := highwayhash.New64(key)
	if err != nil {
		return err
	}
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
		Type:   "HighwayHash-64",
		Key:    hex.EncodeToString(key),
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func HashHighwayHash128(c echo.Context) error {
	var key []byte
	if c.QueryParam("key") == "random" {
		k, err := randKey(32)
		if err != nil {
			return err
		}
		key = k
	} else {
		k, err := hex.DecodeString(c.QueryParam("key"))
		if err != nil {
			e := BasicError{
				Error: "Invalid Hex Value for parameter \"key\"",
			}
			return c.JSON(http.StatusBadRequest, e)
		}
		key = k
	}
	// Check Key Length
	if len(key) != 32 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "HighwayHash key parameter must be 32 bytes",
		})
	}
	hash, err := highwayhash.New128(key)
	if err != nil {
		return err
	}
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
		Type:   "HighwayHash-128",
		Key:    hex.EncodeToString(key),
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}

func randKey(len int) (hexVal []byte, err error) {
	b := make([]byte, len)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	hexVal = b
	return
}
