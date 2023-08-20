package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dwin/hashify/internal/api"
	"github.com/dwin/hashify/internal/config"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("starting.")

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config.")
	}

	log.Info().Str("build", config.AppBuild).Str("app", config.AppName).Msg("starting application.")

	server := api.NewServer(config)

	ctx, cancel := context.WithCancel(context.Background())

	// Listen for SIGINT and SIGTERM signals and cancel the context
	// when they are received.
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigCh)
		<-sigCh
		log.Info().Msg("received signal, stopping server.")
		cancel()
	}()

	log.Info().Str("listen", config.ListenHTTP).Msg("starting server.")

	if err := server.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("server run error.")
	}

	log.Info().Msg("application stopped.")
}
