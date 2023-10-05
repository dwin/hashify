package httpapi_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestAPI_GetKeygenKeyLength(t *testing.T) {
	testServer := LoadTestServer(t)

	expect := httpexpect.Default(t, testServer.URL)

	t.Run("GetKeygenKeyLength-32-OK", func(t *testing.T) {
		resp := expect.GET("/keygen/32").Expect()
		resp.Status(http.StatusOK)
		jsonObj := resp.JSON().Object()
		jsonObj.NotEmpty()
		jsonObj.Keys().ContainsOnly("KeyHex", "Length")
		jsonObj.Value("Length").IsEqual(32)
		jsonObj.Value("KeyHex").String().NotEmpty()
	})

	t.Run("GetKeygenKeyLength-0-BadRequest", func(t *testing.T) {
		resp := expect.GET("/keygen/0").Expect()
		resp.Status(http.StatusBadRequest)
		jsonObj := resp.JSON().Object()
		jsonObj.NotEmpty()
		jsonObj.Keys().ContainsOnly("Error")
		jsonObj.Value("Error").IsEqual("Key length must be between 1 and 256")
	})

	t.Run("GetKeygenKeyLength-1000-BadRequest", func(t *testing.T) {
		resp := expect.GET("/keygen/1000").Expect()
		resp.Status(http.StatusBadRequest)
		jsonObj := resp.JSON().Object()
		jsonObj.NotEmpty()
		jsonObj.Keys().ContainsOnly("Error")
		jsonObj.Value("Error").IsEqual("Key length must be between 1 and 256")
	})
}
