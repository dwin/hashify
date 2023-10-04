package httpapi

import (
	"context"

	"github.com/dwin/hashify/pkg/openapi"
)

func (a *API) GetVersion(ctx context.Context, request openapi.GetVersionRequestObject) (openapi.GetVersionResponseObject, error) {
	return openapi.GetVersion200JSONResponse{
		Version: a.config.AppBuild,
	}, nil
}
