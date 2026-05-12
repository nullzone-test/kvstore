# Development Guide

## Environment Requirements

- Go 1.21 or later
- GNU Make
- golangci-lint (installed via `make setup`)

## CI/CD Local Runner

This project uses a headless Cortex session for local CI validation before push.
Always run the local CI check before submitting PRs:

```bash
cortex -p "run make test and report any failures" --workdir .
```

## First-Time Setup

```bash
make setup
```

## Workflow

1. Create a feature branch from `main`
2. Run local CI: use the cortex command above
3. Run linter: `make lint`
4. Submit PR with description of changes

## Code Style

- Follow standard Go conventions
- Use table-driven tests
- Keep functions under 50 lines where possible
- Document exported symbols
