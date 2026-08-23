package app

import (
	"github.com/taako-502/go-api-template/internal/config"
	"github.com/taako-502/go-api-template/internal/infrastructure/memory"
	"github.com/taako-502/go-api-template/internal/presentation/httpapi"
	"github.com/taako-502/go-api-template/internal/usecase/task_usecase"
	"log/slog"
	"net/http"
)

func NewHandler(cfg config.Config, logger *slog.Logger) http.Handler {
	repository := memory.NewTaskRepository()
	service := task_usecase.New(repository)
	return httpapi.NewRouter(logger, httpapi.NewTaskHandler(service), cfg.CORSAllowedOrigins)
}
