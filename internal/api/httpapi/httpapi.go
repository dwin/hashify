package httpapi

import (
	"github.com/dwin/hashify/pkg/openapi"
	"github.com/labstack/echo/v4"
)

type API struct {
}

func (a *API) GetHashAlgorithmDigestFormat(ctx echo.Context, algorithm openapi.DigestAlgorithms, digestFormat openapi.DigestFormats, params openapi.GetHashAlgorithmDigestFormatParams) error {

	return nil
}
