package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"time"
)

func Run(ctx context.Context, logger *slog.Logger, s *http.Server, timeout time.Duration) error {
	errs := make(chan error, 1)
	go func() { logger.Info("server started", "address", s.Addr); errs <- s.ListenAndServe() }()
	select {
	case err := <-errs:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		shutdown, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		return s.Shutdown(shutdown)
	}
}
func DebugServer(address string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /debug/pprof/", pprof.Index)
	mux.HandleFunc("GET /debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("GET /debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("GET /debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("POST /debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("GET /debug/pprof/trace", pprof.Trace)
	return &http.Server{Addr: address, Handler: mux, ReadHeaderTimeout: 5 * time.Second, MaxHeaderBytes: 1 << 20, MaxHeaderValueCount: 100}
}
