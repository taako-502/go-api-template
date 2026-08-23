package main

import (
	"context"
	"errors"
	"github.com/taako-502/go-api-template/internal/app"
	"github.com/taako-502/go-api-template/internal/config"
	"github.com/taako-502/go-api-template/internal/platform/server"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg, err := config.Load()
	if err != nil {
		logger.Error("load config", "error", err)
		os.Exit(1)
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if cfg.DebugAddress != "" {
		debug := server.DebugServer(cfg.DebugAddress)
		go func() {
			if err := debug.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Error("debug server", "error", err)
			}
		}()
		defer func() {
			if err := debug.Close(); err != nil {
				logger.Error("close debug server", "error", err)
			}
		}()
	}
	httpServer := &http.Server{Addr: cfg.Address, Handler: app.NewHandler(cfg, logger), ReadHeaderTimeout: cfg.ReadHeaderTimeout, IdleTimeout: 60 * time.Second, MaxHeaderBytes: 1 << 20, MaxHeaderValueCount: 100}
	if err := server.Run(ctx, logger, httpServer, cfg.ShutdownTimeout); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
	logger.Info("server stopped")
}
