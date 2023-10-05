package httpapi_test

import (
	"net/http"

	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestAPI_GetVersion(t *testing.T) {
	testServer := LoadTestServer(t)

	expect := httpexpect.Default(t, testServer.URL)

	t.Run("GetVersion-OK", func(t *testing.T) {
		resp := expect.GET("/version").Expect()
		resp.Status(http.StatusOK)
		resp.JSON().Object().Keys().ContainsOnly("version")
	})
}
