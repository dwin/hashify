package httpapi

import (
	"context"

	"github.com/dwin/hashify/pkg/openapi"
)

func (a *API) GetStatus(ctx context.Context, request openapi.GetStatusRequestObject) (openapi.GetStatusResponseObject, error) {
	return openapi.GetStatus200JSONResponse{
		Status:          "OK",
		HashesGenerated: int64(a.metrics.HashCount()),
		KeysGenerated:   int64(a.metrics.KeyGenCount()),
		Uptime:          a.metrics.Uptime().String(),
	}, nil
}
