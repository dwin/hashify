package hashers

import (
	"github.com/dwin/hashify/internal/hasher"
	"github.com/dwin/hashify/pkg/openapi"
)

type RequiredKeyLength int

const (
	NoRequiredKeyLength RequiredKeyLength = 0
)

var Hashers = map[openapi.DigestAlgorithms]hasher.Hasher{
	openapi.Blake2b256: Blake2b256{},
}
