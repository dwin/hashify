package httpapi_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestAPI_GetMethods(t *testing.T) {
	testServer := LoadTestServer(t)

	expect := httpexpect.Default(t, testServer.URL)

	t.Run("GetMethods-OK", func(t *testing.T) {
		resp := expect.GET("/methods").Expect()
		resp.Status(http.StatusOK)
		jsonArr := resp.JSON().Array()
		jsonArr.NotEmpty()
		for _, val := range jsonArr.Iter() {
			obj := val.Object()
			obj.Keys().ContainsOnly("Endpoint", "MaxKeyLength", "MinKeyLength", "Name")
		}
	})
}
