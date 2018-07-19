package controller

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

var StartTime time.Time
var hashesGenerated int
var keysGenerated int

func GetStatus(c echo.Context) error {
	name, err := os.Hostname()
	if err != nil {
		log.Println("Unable to get Hostname for status, error: ", err)
	}
	j, err := json.Marshal(map[string]interface{}{
		"status":          "OK",
		"uptime":          time.Since(StartTime).String(),
		"hostname":        name,
		"hashesGenerated": hashesGenerated,
		"keysGenerated":   keysGenerated,
	})
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
