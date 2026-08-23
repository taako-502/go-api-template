package main

import (
	"encoding/json/jsontext"
	jsonv2 "encoding/json/v2"
	"flag"
	"fmt"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/taako-502/go-api-template/internal/infrastructure/memory"
	"github.com/taako-502/go-api-template/internal/presentation/httpapi"
	"github.com/taako-502/go-api-template/internal/usecase/task_usecase"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	output := flag.String("output", "docs/openapi.json", "output path")
	flag.Parse()
	mux := http.NewServeMux()
	api := humago.New(mux, httpapi.NewHumaConfig())
	httpapi.RegisterOpenAPIRoutes(api, httpapi.NewTaskHandler(task_usecase.New(memory.NewTaskRepository())))
	payload, err := jsonv2.Marshal(api.OpenAPI(), jsontext.WithIndent("  "), jsonv2.Deterministic(true))
	if err != nil {
		fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(*output), 0o755); err != nil {
		fatal(err)
	}
	if err := os.WriteFile(*output, payload, 0o644); err != nil {
		fatal(err)
	}
}
func fatal(err error) { _, _ = fmt.Fprintln(os.Stderr, err); os.Exit(1) }
