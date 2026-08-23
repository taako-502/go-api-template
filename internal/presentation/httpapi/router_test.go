package httpapi_test

import (
	"bytes"
	jsonv2 "encoding/json/v2"
	"github.com/taako-502/go-api-template/internal/app"
	"github.com/taako-502/go-api-template/internal/config"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskFlow(t *testing.T) {
	s := httptest.NewServer(app.NewHandler(config.Config{CORSAllowedOrigins: []string{"http://localhost:3000"}}, slog.New(slog.NewTextHandler(io.Discard, nil))))
	t.Cleanup(s.Close)
	response, err := http.Post(s.URL+"/v1/tasks", "application/json", bytes.NewBufferString(`{"title":"write tests"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("status=%d", response.StatusCode)
	}
	var created struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}
	if err := jsonv2.UnmarshalRead(response.Body, &created); err != nil {
		t.Fatal(err)
	}
	if created.ID == "" || created.Title != "write tests" {
		t.Fatalf("created=%+v", created)
	}
	get, err := http.Get(s.URL + "/v1/tasks/" + created.ID)
	if err != nil {
		t.Fatal(err)
	}
	defer get.Body.Close()
	if get.StatusCode != http.StatusOK {
		t.Fatalf("status=%d", get.StatusCode)
	}
}
