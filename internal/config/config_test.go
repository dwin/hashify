package config_test

import (
	"testing"

	"github.com/dwin/hashify/internal/config"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestConfig_ConfigureLogger(t *testing.T) {
	t.Setenv("LOG_LEVEL", "trace")

	cfg, err := config.LoadConfig()
	require.NoError(t, err)

	cfg.ConfigureLogger()

	require.Equal(t, zerolog.TraceLevel, cfg.LogLevel)
	require.Equal(t, zerolog.TraceLevel, zerolog.GlobalLevel())
}
