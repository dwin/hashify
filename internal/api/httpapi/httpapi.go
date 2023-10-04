package httpapi

import (
	"fmt"
	"net/http"

	"github.com/dwin/hashify/internal/config"
	"github.com/dwin/hashify/internal/metrics"
	"github.com/dwin/hashify/pkg/openapi"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type API struct {
	openapi.StrictServerInterface
	config  *config.Config
	metrics *metrics.Collector
}

func NewHTTPAPI(c *config.Config, metrics *metrics.Collector) *API {
	return &API{
		config:  c,
		metrics: metrics,
	}
}

func (a *API) Load() (http.Handler, error) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("failed to load swagger: %w", err)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create Echo Router instance
	echoRouter := echo.New()

	// Setup Middlewares

	// Log all requests
	echoRouter.Use(echomiddleware.RequestLoggerWithConfig(echomiddleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogError:    true,
		LogMethod:   true,
		LogLatency:  true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v echomiddleware.RequestLoggerValues) error {
			logger := log.With().
				Str("http.method", v.Method).
				Str("http.url", v.URI).
				Int("http.status_code", v.Status).
				Str("http.user_agent", v.UserAgent).
				Dur("duration", v.Latency).
				Logger()

			if v.Error == nil {
				logger.Info().Msg("http request.")
			} else {
				logger.Error().
					Err(v.Error).
					Msg("http request with error.")
			}
			return nil
		},
	}))

	strictHandler := openapi.NewStrictHandler(a, nil)

	openapi.RegisterHandlers(echoRouter, strictHandler)

	return echoRouter, nil
}
