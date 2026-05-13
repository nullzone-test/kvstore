# kvstore — Development

## Prerequisites

- Go 1.21+
- GNU Make
- golangci-lint (`make setup` installs it)

## Setup

```bash
make setup
make build
make test
```

## Usage

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

## Roadmap

- [ ] SSTable reader with binary search
- [ ] Multi-level compaction
- [ ] Snapshot isolation
- [ ] HTTP API layer
- [ ] Prometheus metrics
