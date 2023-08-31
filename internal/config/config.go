package config

import (
	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppBuild   string `env:"APP_BUILD"`
	ListenHTTP string `env:"LISTEN_HTTP" envDefault:"127.0.0.1:32005"`
}

func LoadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
