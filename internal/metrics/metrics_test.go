package metrics_test

import (
	"testing"

	"github.com/dwin/hashify/internal/metrics"
	"github.com/stretchr/testify/require"
)

func TestCollector(t *testing.T) {
	collector := metrics.NewCollector()
	require.NotNil(t, collector)

	collector.KeyGenerations(1)
	collector.HashOperations("sha256", "hex")

	require.Equal(t, uint64(1), collector.HashCount())
	require.Equal(t, uint64(1), collector.KeyGenCount())
	require.NotZero(t, collector.Uptime())
}
