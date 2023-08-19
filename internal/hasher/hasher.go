package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"

	"github.com/dwin/hashify/pkg/openapi"
	"github.com/minio/highwayhash"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"

	"github.com/valyala/bytebufferpool"
)

var ErrUnsupportedAlgorithmFormat = errors.New("unsupported algorithm or format")

func Hash(algo openapi.DigestAlgorithms, input io.Reader, key ...byte) ([]byte, error) {
	var hasher hash.Hash
	var err error

	switch algo {
	case openapi.Blake2b256:
		hasher, err = blake2b.New256(key)
	case openapi.Blake2b384:
		hasher, err = blake2b.New384(key)
	case openapi.Blake2b512:
		hasher, err = blake2b.New512(key)
	case openapi.Blake2s128:
		hasher, err = blake2s.New128(key)
	case openapi.Blake2s256:
		hasher, err = blake2s.New256(key)
	case openapi.HighwayHash64:
		hasher, err = highwayhash.New64(key)
	case openapi.HighwayHash128:
		hasher, err = highwayhash.New128(key)
	case openapi.HighwayHash256:
		hasher, err = highwayhash.New(key)
	case openapi.MD4:
		hasher = md4.New()
	case openapi.MD5:
		hasher = md5.New()
	case openapi.SHA1:
		hasher = sha1.New()
	case openapi.SHA256:
		hasher = sha256.New()
	case openapi.SHA384:
		hasher = sha512.New384()
	case openapi.SHA512:
		hasher = sha512.New()
	case openapi.SHA512256:
		hasher = sha512.New512_256()
	case openapi.SHA3256:
		hasher = sha3.New256()
	case openapi.SHA3384:
		hasher = sha3.New384()
	case openapi.SHA3512:
		hasher = sha3.New512()
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedAlgorithmFormat, algo)
	}

	if err != nil {
		return nil, fmt.Errorf("error initializing hasher: %w", err)
	}

	if _, err := io.Copy(hasher, input); err != nil {
		return nil, fmt.Errorf("error hashing input: %w", err)
	}

	return hasher.Sum(nil), nil
}

func GetDigest(format openapi.DigestFormats, hash []byte) (string, error) {
	bb := bytebufferpool.Get()

	var err error

	switch format {
	case openapi.Base32:
		_, err = base32.NewEncoder(base32.StdEncoding, bb).Write(hash)
	case openapi.Base64:
		_, err = base64.NewEncoder(base64.StdEncoding, bb).Write(hash)
	case openapi.Base64url:
		_, err = base64.NewEncoder(base64.URLEncoding, bb).Write(hash)
	case openapi.Hex:
		_, err = hex.NewEncoder(bb).Write(hash)
	default:
		return "", fmt.Errorf("%w: %s", ErrUnsupportedAlgorithmFormat, format)
	}

	if err != nil {
		return "", fmt.Errorf("error encoding hash: %w", err)
	}

	defer func() {
		bytebufferpool.Put(bb)
	}()

	return bb.String(), nil
}
