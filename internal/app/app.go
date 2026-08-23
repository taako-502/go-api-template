package app

import (
	"github.com/taako-502/go-api-template/internal/application/taskservice"
	"github.com/taako-502/go-api-template/internal/config"
	"github.com/taako-502/go-api-template/internal/infrastructure/memory"
	"github.com/taako-502/go-api-template/internal/presentation/httpapi"
	"log/slog"
	"net/http"
)

func NewHandler(cfg config.Config, logger *slog.Logger) http.Handler {
	repository := memory.NewTaskRepository()
	service := taskservice.New(repository)
	return httpapi.NewRouter(logger, httpapi.NewTaskHandler(service), cfg.CORSAllowedOrigins)
}
