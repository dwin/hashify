package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var StartTime time.Time

func GetStatus(c echo.Context) error {
	j, err := json.Marshal(map[string]string{
		"status": "OK",
		"uptime": time.Since(StartTime).String(),
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
