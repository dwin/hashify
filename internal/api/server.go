package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/dwin/hashify/internal/api/httpapi"
	"github.com/dwin/hashify/internal/config"
	"github.com/dwin/hashify/internal/metrics"
)

const serverTimeout = 10 * time.Second

type Server struct {
	config  *config.Config
	metrics *metrics.Collector
}

func NewServer(c *config.Config, metrics *metrics.Collector) *Server {
	return &Server{
		config:  c,
		metrics: metrics,
	}
}

func (s *Server) Start(ctx context.Context) error {
	httpAPIHandler, err := httpapi.NewHTTPAPI(s.config, s.metrics).Load()
	if err != nil {
		return fmt.Errorf("failed to load httpapi: %w", err)
	}

	// Create HTTP Server
	httpServer := http.Server{
		Addr:              s.config.ListenHTTP,
		Handler:           httpAPIHandler,
		ReadHeaderTimeout: serverTimeout,
		ReadTimeout:       serverTimeout,
		WriteTimeout:      serverTimeout,
	}

	go func() {
		<-ctx.Done()
		httpServer.Shutdown(ctx)
	}()

	log.Info().Msgf("Starting HTTP Server on '%s'", s.config.ListenHTTP)

	// Start HTTP Server
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http server error: %w", err)
	}

	return nil
}
