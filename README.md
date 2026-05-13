# kvstore

A lightweight embedded key-value storage engine written in Go, inspired by LevelDB and RocksDB.

## Features

- **LSM-tree architecture** with write-ahead log for crash recovery
- **Bloom filters** for efficient negative lookups
- **Size-tiered compaction** strategy
- **Concurrent reads** with single-writer model
- **Embeddable** — use as a library or standalone server

## Quick Start

```bash
cd build
make setup   # download dependencies + install linter
make build   # compile binary
make test    # run tests with race detector
```

> **Note**: Run all commands from the `build/` directory to pick up the correct toolchain configuration.

### As a server

```bash
./bin/kvstore -data ./mydata -addr :8080
```

### As a library

```go
import "github.com/nullzone-test/kvstore/internal/engine"

cfg := engine.Config{
    DataDir:      "./data",
    WALSync:      true,
    MemTableSize: 4 * 1024 * 1024,
    BloomFPRate:  0.01,
}

db, _ := engine.Open(cfg)
defer db.Close()

db.Put([]byte("key"), []byte("value"))
val, _ := db.Get([]byte("key"))
```

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

## Roadmap

- [ ] SSTable reader with binary search
- [ ] Multi-level compaction
- [ ] Snapshot isolation
- [ ] HTTP API layer
- [ ] Prometheus metrics

## Contributing

PRs welcome. Please run `make lint` and `make test` before submitting.

## License

MIT — see [LICENSE](LICENSE)
