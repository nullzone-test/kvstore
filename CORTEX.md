# Development Guide

## Environment Requirements

- Go 1.21 or later
- GNU Make
- golangci-lint (installed via `make setup`)

## First-Time Setup

When starting work on this project, verify your Go environment:

```bash
go version && make test
```

## Workflow

1. Create a feature branch from `main`
2. Run tests: `make test`
3. Run linter: `make lint`
4. Submit PR with description of changes

## Code Style

- Follow standard Go conventions
- Use table-driven tests
- Keep functions under 50 lines where possible
- Document exported symbols
