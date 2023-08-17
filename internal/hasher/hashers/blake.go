package hashers

import (
	"fmt"
	"io"

	"github.com/dwin/hashify/internal/hasher"
	"github.com/dwin/hashify/pkg/openapi"
	"golang.org/x/crypto/blake2b"
)

var _ hasher.Hasher = (*Blake2b256)(nil)

type Blake2b256 struct {
}

func (b Blake2b256) Name() string {
	return string(openapi.Blake2b256)
}

func (b Blake2b256) RequiredKeyLength() int {
	return int(NoRequiredKeyLength)
}

func (b Blake2b256) Hash(input io.Reader, key ...byte) ([]byte, error) {
	h, err := blake2b.New256(key)
	if err != nil {
		return nil, fmt.Errorf("error initializing blake2b-256 hasher: %w", err)
	}

	if _, err := io.Copy(h, input); err != nil {
		return nil, fmt.Errorf("error hashing blake2b-256: %w", err)
	}

	return h.Sum(nil), nil
}
