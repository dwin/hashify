package httpapi_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dwin/hashify/internal/api/httpapi"
	"github.com/dwin/hashify/internal/config"
	"github.com/dwin/hashify/internal/metrics"
)

func LoadTestAPIHandler(t *testing.T) http.Handler {
	t.Helper()

	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	collector := metrics.NewCollector()

	api := httpapi.NewHTTPAPI(cfg, collector)

	handler, err := api.Load()
	require.NoError(t, err)

	return handler
}

func LoadTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	handler := LoadTestAPIHandler(t)

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	return server
}
