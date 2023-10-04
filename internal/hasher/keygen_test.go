package hasher_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/dwin/hashify/internal/hasher"
	"github.com/stretchr/testify/require"
)

func TestRandomKeyHex(t *testing.T) {
	lengths := []int{16, 32, 64, 128, 256, 512, 1024}

	for _, length := range lengths {
		t.Run(fmt.Sprintf("length=%d", length), func(t *testing.T) {
			key, err := hasher.RandomKeyHex(length)
			require.NoError(t, err)

			decodedKey, err := hex.DecodeString(key)
			if err != nil {
				t.Errorf("failed to decode key: %v", err)
			}

			require.Len(t, decodedKey, length)
		})
	}
}
