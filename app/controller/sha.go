package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"

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
