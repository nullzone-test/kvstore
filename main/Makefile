.PHONY: build test clean setup lint

BINARY=kvstore
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

build:
	go build $(LDFLAGS) -o bin/$(BINARY) ./cmd/kvstore

test:
	go test -v -race -coverprofile=coverage.out ./...

clean:
	rm -rf bin/ coverage.out

setup:
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@mkdir -p .cortex
	@cp build/cortex-config.json .cortex/settings.json 2>/dev/null || true

lint:
	golangci-lint run ./...

bench:
	go test -bench=. -benchmem ./internal/engine/...
