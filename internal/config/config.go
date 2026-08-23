package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	Environment        string
	Address            string
	DebugAddress       string
	CORSAllowedOrigins []string
	ReadHeaderTimeout  time.Duration
	ShutdownTimeout    time.Duration
}

func Load() (Config, error) {
	cfg := Config{Environment: value("APP_ENV", "local"), Address: value("APP_ADDR", ":8080"), DebugAddress: strings.TrimSpace(os.Getenv("DEBUG_ADDR")), CORSAllowedOrigins: list("CORS_ALLOW_ORIGINS", []string{"http://localhost:3000"}), ReadHeaderTimeout: 5 * time.Second, ShutdownTimeout: 10 * time.Second}
	if !strings.Contains(cfg.Address, ":") {
		return Config{}, fmt.Errorf("APP_ADDR must be a host:port address")
	}
	if cfg.DebugAddress != "" && !isLoopbackAddress(cfg.DebugAddress) {
		return Config{}, fmt.Errorf("DEBUG_ADDR must bind to loopback")
	}
	return cfg, nil
}

func value(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}
func list(key string, fallback []string) []string {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback
	}
	parts := strings.Split(raw, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if v := strings.TrimSpace(part); v != "" {
			result = append(result, v)
		}
	}
	return result
}
func isLoopbackAddress(addr string) bool {
	return strings.HasPrefix(addr, "127.0.0.1:") || strings.HasPrefix(addr, "localhost:") || strings.HasPrefix(addr, "[::1]:")
}
