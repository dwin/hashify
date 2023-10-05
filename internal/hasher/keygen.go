package hasher

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func RandomKeyHex(length int) (string, error) {
	b := make([]byte, length)

	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("rand read error: %w", err)
	}

	return hex.EncodeToString(b), nil
}
