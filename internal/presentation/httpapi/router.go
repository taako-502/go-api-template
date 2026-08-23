package httpapi

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"log/slog"
	"net/http"
	"slices"
	"strings"
	"time"
	"uuid"
)

func NewRouter(logger *slog.Logger, h TaskHandler, allowed []string) http.Handler {
	mux := http.NewServeMux()
	api := humago.New(mux, NewHumaConfig())
	RegisterOpenAPIRoutes(api, h)
	mux.HandleFunc("GET /health/live", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) })
	mux.HandleFunc("GET /health/ready", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) })
	return recoverPanic(logger, accessLog(logger, requestID(cors(mux, allowed))))
}
func NewHumaConfig() huma.Config {
	c := huma.DefaultConfig("Go API Template", "1.0.0")
	c.DocsPath = "/docs"
	c.OpenAPIPath = "/openapi.json"
	return c
}
func RegisterOpenAPIRoutes(api huma.API, h TaskHandler) {
	huma.Register(api, huma.Operation{OperationID: "create-task", Method: http.MethodPost, Path: "/v1/tasks", Tags: []string{"tasks"}, DefaultStatus: http.StatusCreated}, h.Create)
	huma.Register(api, huma.Operation{OperationID: "list-tasks", Method: http.MethodGet, Path: "/v1/tasks", Tags: []string{"tasks"}}, h.List)
	huma.Register(api, huma.Operation{OperationID: "get-task", Method: http.MethodGet, Path: "/v1/tasks/{id}", Tags: []string{"tasks"}}, h.Get)
	huma.Register(api, huma.Operation{OperationID: "update-task", Method: http.MethodPut, Path: "/v1/tasks/{id}", Tags: []string{"tasks"}}, h.Update)
	huma.Register(api, huma.Operation{OperationID: "delete-task", Method: http.MethodDelete, Path: "/v1/tasks/{id}", Tags: []string{"tasks"}, DefaultStatus: http.StatusNoContent}, h.Delete)
}
func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSpace(r.Header.Get("X-Request-ID"))
		if id == "" {
			id = uuid.New().String()
		}
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	})
}
func cors(next http.Handler, allowed []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && slices.Contains(allowed, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Add("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Request-ID")
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func recoverPanic(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if v := recover(); v != nil {
				logger.ErrorContext(r.Context(), "panic recovered", "panic", v)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
func accessLog(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		next.ServeHTTP(w, r)
		logger.InfoContext(r.Context(), "request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(started))
	})
}
