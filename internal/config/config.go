package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/carlmjohnson/versioninfo"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/rs/zerolog"
)

type Config struct {
	AppName           string        `env:"APP_NAME"`
	AppBuild          string        `env:"APP_BUILD"`
	LogLevel          zerolog.Level `env:"LOG_LEVEL" envDefault:"info"`
	ListenHTTP        string        `env:"LISTEN_HTTP" envDefault:"127.0.0.1:32005"`
	ListenHTTPMetrics string        `env:"LISTEN_HTTP_METRICS" envDefault:"127.0.0.1:32006"`
}

func LoadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	if cfg.AppBuild == "" {
		cfg.AppBuild = versioninfo.Short()
	}

	return &cfg, nil
}

func (c Config) ConfigureLogger() {
	zerolog.SetGlobalLevel(c.LogLevel)
}
