package hasher

import (
	"io"
)

//go:generate mockery --name Hasher --filename=hasher.go --outpkg=hashermocks --output hashermocks
type Hasher interface {
	Name() string
	RequiredKeyLength() int
	Hash(input io.Reader, key ...byte) ([]byte, error)
}
