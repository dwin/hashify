package controller

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

func HashMD5(c echo.Context) error {
	hash := md5.New()
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
		Type:   "MD5",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
