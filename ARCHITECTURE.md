# Architecture

Dependency direction is always inward:

`cmd -> presentation/infrastructure -> application -> domain`

- `internal/domain`: entities, domain errors, repository contracts; no framework dependency.
- `internal/application`: use cases; depends only on domain.
- `internal/infrastructure`: replaceable outbound adapters; the sample uses memory.
- `internal/presentation`: Huma + `net/http` inbound adapter and middleware.
- `internal/app`: composition root shared by the executable and integration tests.
- `cmd/api`: process lifecycle and dependency wiring.

Add PostgreSQL, Redis, authentication, or cloud SDKs as infrastructure adapters only when required. Do not expose storage models as API models.
