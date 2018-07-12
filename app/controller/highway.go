package controller

import (
	"encoding/hex"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/minio/highwayhash"
)

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

// TODO: Obtain header key value does not work
func parseHighwayHashKey(c echo.Context) (key []byte, err error) {
	if c.Request().Method == http.MethodPost {
		fmt.Println("parsing post")
		if c.Request().Header.Get("X-Hashify-Key") == "random" {
			// Generate random key if Header passed with value "random"
			k, err := randKey(32)
			if err != nil {
				return nil, err
			}
			return k, err
		} else {
			k, err := hex.DecodeString(c.Request().Header.Get("X-Hashify-Key"))
			if err != nil {
				return nil, err
			}
			return k, err
		}
	}
	// If GET
	fmt.Println("parsing query")
	if c.QueryParam("key") == "random" {
		k, err := randKey(32)
		if err != nil {
			return nil, err
		}
		key = k

	} else {
		k, err := hex.DecodeString(c.QueryParam("key"))
		if err != nil {
			return nil, err
		}
		key = k

	}
	return
}
