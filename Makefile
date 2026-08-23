.PHONY: run test test-race tidy fix fmt vet lint vuln verify build openapi check

GOLANGCI_LINT_VERSION ?= v2.13.1
GOLANGCI_LINT_CACHE ?= $(CURDIR)/.cache/golangci-lint
GOVULNCHECK_VERSION ?= v1.7.0

run:
	go run ./cmd/api
test:
	go test ./...
test-race:
	go test -race ./...
tidy:
	go mod tidy
fix:
	go fix ./...
fmt:
	gofmt -w -s $$(git ls-files '*.go')
vet:
	go vet ./...
lint:
	GOLANGCI_LINT_CACHE=$(GOLANGCI_LINT_CACHE) \
		go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./...
vuln:
	go run golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION) ./...
verify:
	go mod download
	go mod verify
build:
	go build -trimpath -ldflags='-s -w' -o bin/api ./cmd/api
openapi:
	go run ./cmd/openapi -output docs/api/openapi.json
check: tidy fix fmt vet openapi
