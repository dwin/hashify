package httpapi_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestAPI_GetStatus(t *testing.T) {
	testServer := LoadTestServer(t)

	expect := httpexpect.Default(t, testServer.URL)

	t.Run("GetStatus-OK", func(t *testing.T) {
		resp := expect.GET("/status").Expect()
		resp.Status(http.StatusOK)
		jsonObj := resp.JSON().Object()
		jsonObj.Keys().ContainsOnly("status", "hashesGenerated", "keysGenerated", "uptime")
		jsonObj.Value("status").IsEqual("OK")
		jsonObj.Value("hashesGenerated").IsNumber()
		jsonObj.Value("keysGenerated").IsNumber()
		jsonObj.Value("uptime").IsString()
	})
}
