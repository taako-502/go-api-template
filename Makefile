.PHONY: run test test-race lint vuln build openapi check

run:
	go run ./cmd/api
test:
	go test ./...
test-race:
	go test -race ./...
lint:
	go vet ./...
vuln:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
build:
	go build -trimpath -ldflags='-s -w' -o bin/api ./cmd/api
openapi:
	go run ./cmd/openapi -output docs/openapi.json
check: lint test-race
