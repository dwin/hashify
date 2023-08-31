package httpapi

import (
	"context"

	"github.com/dwin/hashify/pkg/openapi"
)

func (a *API) GetMethods(ctx context.Context, request openapi.GetMethodsRequestObject) (openapi.GetMethodsResponseObject, error) {
	return openapi.GetMethods200JSONResponse{
		Body: []openapi.HashAlgorithm{
			{Name: openapi.Blake2b256, Endpoint: "/hash/BLAKE2B-256", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.Blake2b384, Endpoint: "/hash/BLAKE2B-384", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.Blake2b512, Endpoint: "/hash/BLAKE2B-512", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.Blake2s128, Endpoint: "/hash/BLAKE2s-128", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.Blake2s256, Endpoint: "/hash/BLAKE2s-256", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.HighwayHash256, Endpoint: "/hash/HIGHWAY", MinKeyLength: 32, MaxKeyLength: 32},
			{Name: openapi.HighwayHash64, Endpoint: "/hash/HIGHWAY-64", MinKeyLength: 32, MaxKeyLength: 32},
			{Name: openapi.HighwayHash128, Endpoint: "/hash/HIGHWAY-128", MinKeyLength: 32, MaxKeyLength: 32},
			{Name: openapi.MD4, Endpoint: "/hash/MD4", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.MD5, Endpoint: "/hash/MD5", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.SHA1, Endpoint: "/hash/SHA1", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.SHA256, Endpoint: "/hash/SHA256", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.SHA384, Endpoint: "/hash/SHA384", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.SHA512, Endpoint: "/hash/SHA512", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.SHA512256, Endpoint: "/hash/SHA512-256", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.SHA3256, Endpoint: "/hash/SHA3-256", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.SHA3384, Endpoint: "/hash/SHA3-384", MinKeyLength: 0, MaxKeyLength: 0},
			{Name: openapi.SHA3512, Endpoint: "/hash/SHA3-512", MinKeyLength: 0, MaxKeyLength: 0},
		},
	}, nil
}
