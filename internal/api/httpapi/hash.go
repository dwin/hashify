package httpapi

import (
	"context"

	"github.com/dwin/hashify/pkg/openapi"
)

func (a *API) GetHashAlgorithmDigestFormat(ctx context.Context, request openapi.GetHashAlgorithmDigestFormatRequestObject) (openapi.GetHashAlgorithmDigestFormatResponseObject, error) {

	return openapi.GetHashAlgorithmDigestFormat200JSONResponse{
		Digest: "digest",
	}, nil
}
