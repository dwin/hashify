package httpapi

import (
	"context"

	"github.com/dwin/hashify/internal/hasher"
	"github.com/dwin/hashify/pkg/openapi"
)

func (a *API) GetKeygenKeyLength(ctx context.Context, request openapi.GetKeygenKeyLengthRequestObject) (openapi.GetKeygenKeyLengthResponseObject, error) {
	keyLength := request.KeyLength

	if keyLength < 1 || keyLength > 256 {
		msg := "Key length must be between 1 and 256"
		return openapi.GetKeygenKeyLength400JSONResponse{
			Error: &msg,
		}, nil
	}

	defer a.metrics.KeyGenerations(keyLength)

	hexKey, err := hasher.RandomKeyHex(keyLength)
	if err != nil {
		msg := "Error generating key"
		return openapi.GetKeygenKeyLength500JSONResponse{
			Error: &msg,
		}, err
	}

	return openapi.GetKeygenKeyLength200JSONResponse{
		KeyHex: hexKey,
		Length: keyLength,
	}, nil
}
