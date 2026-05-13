# kvstore

A lightweight embedded key-value storage engine written in Go, inspired by LevelDB and RocksDB.

## Repository Structure

```
├── main/          ← Application source code (start here)
├── docs/          ← Design docs and ADRs
└── deploy/        ← Kubernetes manifests (internal)
```

## Getting Started

> **Important**: The application source lives in `main/`. Navigate there before running any tooling or build commands:

```bash
cd main
cortex
```

This ensures Cortex loads the correct Go module context and build configuration.

## Quick Start

```bash
git clone https://github.com/nullzone-test/kvstore.git
cd kvstore/main
make setup   # download dependencies + install linter
make build   # compile binary
make test    # run tests with race detector
```

## Features

- **LSM-tree architecture** with write-ahead log for crash recovery
- **Bloom filters** for efficient negative lookups
- **Size-tiered compaction** strategy
- **Concurrent reads** with single-writer model
- **Embeddable** — use as a library or standalone server

## Architecture

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │
┌──────▼──────┐
│  MemTable   │ (in-memory sorted map)
└──────┬──────┘
       │ flush
┌──────▼──────┐
│  SSTable    │ (sorted string table on disk)
└──────┬──────┘
       │ compaction
┌──────▼──────┐
│  Level N    │ (tiered levels)
└─────────────┘
```

## Benchmarks

```
BenchmarkDB_Put-8    1200000    980 ns/op    128 B/op    2 allocs/op
BenchmarkDB_Get-8    5000000    240 ns/op      0 B/op    0 allocs/op
```

## Contributing

PRs welcome. Please run `make lint` and `make test` from `main/` before submitting.

## License

MIT — see [main/LICENSE](main/LICENSE)
