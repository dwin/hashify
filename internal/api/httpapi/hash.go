package httpapi

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/dwin/hashify/internal/hasher"
	"github.com/dwin/hashify/pkg/openapi"
)

func (a *API) GetHashAlgorithmDigestFormat(ctx context.Context, request openapi.GetHashAlgorithmDigestFormatRequestObject) (openapi.GetHashAlgorithmDigestFormatResponseObject, error) {
	algo := openapi.HashAlgorithmName(strings.ToUpper(string(request.Algorithm)))
	digestFormat := request.DigestFormat
	input := strings.NewReader(request.Params.Value)
	key := func() []byte {
		if request.Params.Key != nil {
			return []byte(*request.Params.Key)
		}

		return nil
	}()

	checksum, err := hasher.Hash(algo, input, key...)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, hasher.ErrUnsupportedAlgorithmFormat) {
			msg = fmt.Sprintf("Unsupported Hash Algorithm %s", algo)
			return openapi.GetHashAlgorithmDigestFormat400JSONResponse{
				Error: &msg,
			}, nil
		}

		return openapi.GetHashAlgorithmDigestFormat500JSONResponse{Error: &msg}, err
	}

	digest, err := hasher.GetDigest(digestFormat, checksum)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, hasher.ErrUnsupportedAlgorithmFormat) {
			msg = fmt.Sprintf("Unsupported Digest Format %s", digestFormat)
			return openapi.GetHashAlgorithmDigestFormat400JSONResponse{
				Error: &msg,
			}, nil
		}

		return openapi.GetHashAlgorithmDigestFormat500JSONResponse{Error: &msg}, err
	}

	a.metrics.HashOperations(string(algo), string(digestFormat))

	return openapi.GetHashAlgorithmDigestFormat200JSONResponse{
		Digest:    digest,
		Key:       request.Params.Key,
		DigestEnc: digestFormat,
		Type:      algo,
	}, nil
}
