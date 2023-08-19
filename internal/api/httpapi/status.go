package httpapi

import (
	"context"

	"github.com/dwin/hashify/pkg/openapi"
)

func (a *API) GetStatus(ctx context.Context, request openapi.GetStatusRequestObject) (openapi.GetStatusResponseObject, error) {

	return openapi.GetStatus200JSONResponse{
		Status: "OK",
	}, nil
}
