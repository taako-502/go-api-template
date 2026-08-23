# go-api-template

Go 1.27 API server template using [Huma](https://huma.rocks/) on the standard `net/http` router. The layering comes from `nakith-api`; application-specific dependencies were removed.

## Included

- Huma endpoints, validation, Swagger UI, and OpenAPI 3.1 generation
- Domain / application / infrastructure / presentation separation
- Replaceable in-memory repository and complete task CRUD example
- `slog`, request IDs, panic recovery, strict CORS allow-list
- Health endpoints, graceful shutdown, race-test CI, vet, vulnerability scan
- Optional loopback-only pprof server with Go 1.27 `goroutineleak`

## Go 1.27 usage

- Generic methods: typed `Page[T].Map[U]`
- Standard `uuid` package with sortable UUIDv7 identifiers
- `encoding/json/v2` for deterministic OpenAPI and tests
- `http.Server.MaxHeaderValueCount` hardening
- Runtime goroutine-leak profile through pprof

Experimental SIMD is excluded because it has no API-server use case and would make normal builds experimental.

## Run

```sh
cp .env.example .env
set -a; . ./.env; set +a
go mod tidy
make run
```

- API: `http://localhost:8080`
- Swagger UI: `http://localhost:8080/docs`
- OpenAPI: `http://localhost:8080/openapi.json`
- Health: `/health/live`, `/health/ready`
- pprof when enabled: `http://127.0.0.1:6060/debug/pprof/`

```sh
curl -X POST http://localhost:8080/v1/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"replace the sample feature"}'
```

Read [ARCHITECTURE.md](ARCHITECTURE.md) before replacing the sample feature.
