package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

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

	config.ConfigureLogger()

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
		log.Info().Msg("received signal, stopping application.")
		cancel()
	}()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := metrics.RunHTTPMetricsServer(ctx, config.ListenHTTPMetrics); err != nil {
			return fmt.Errorf("metrics server run error: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		if err := server.Start(ctx); err != nil {
			return fmt.Errorf("server run error: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal().Err(err).Msg("error running servers.")
	}

	log.Info().Msg("application stopped.")
}
