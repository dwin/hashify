package httpapi

import (
	"fmt"
	"net/http"

	"github.com/dwin/hashify/internal/config"
	"github.com/dwin/hashify/pkg/openapi"
	"github.com/labstack/echo/v4"
)

type API struct {
	openapi.StrictServerInterface
	config *config.Config
}

func NewHTTPAPI(c *config.Config) *API {
	return &API{
		config: c,
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
	// e.Use(echomiddleware.Logger())

	strictHandler := openapi.NewStrictHandler(a, nil)
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	// echoRouter.Use(middleware.OapiRequestValidator(swagger))

	openapi.RegisterHandlers(echoRouter, strictHandler)

	return echoRouter, nil
}
