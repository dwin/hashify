package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetStatus(c echo.Context) error {
	j, err := json.Marshal(map[string]string{
		"status": "OK",
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
