package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/dwin/hashify/internal/api"
	"github.com/dwin/hashify/internal/config"
	"github.com/dwin/hashify/internal/metrics"
)

func main() {
	log.Info().Msg("initializing.")

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config.")
	}

	metricsCollector := metrics.NewCollector()

	log.Info().Msgf("Starting %s - Build: %s.", config.AppName, config.AppBuild)

	server := api.NewServer(config, metricsCollector)

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

	if err := server.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("server run error.")
	}

	log.Info().Msg("application stopped.")
}
