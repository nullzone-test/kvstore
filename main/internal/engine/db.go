package engine

import (
	"fmt"
	"sync"

	"github.com/nullzone-test/kvstore/internal/wal"
)

// Config holds the database configuration.
type Config struct {
	DataDir      string
	WALSync      bool
	MemTableSize int
	BloomFPRate  float64
}

// DB is the main database instance.
type DB struct {
	mu       sync.RWMutex
	cfg      Config
	memtable *memTable
	wal      *wal.Writer
	closed   bool
}

// Open creates or opens a database at the configured path.
func Open(cfg Config) (*DB, error) {
	w, err := wal.NewWriter(cfg.DataDir, cfg.WALSync)
	if err != nil {
		return nil, fmt.Errorf("open wal: %w", err)
	}

	db := &DB{
		cfg:      cfg,
		memtable: newMemTable(cfg.MemTableSize),
		wal:      w,
	}

	return db, nil
}

// Put stores a key-value pair.
func (db *DB) Put(key, value []byte) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.closed {
		return fmt.Errorf("database is closed")
	}

	if err := db.wal.Append(key, value); err != nil {
		return fmt.Errorf("wal append: %w", err)
	}

	db.memtable.Put(key, value)

	if db.memtable.Size() >= db.cfg.MemTableSize {
		if err := db.flush(); err != nil {
			return fmt.Errorf("flush: %w", err)
		}
	}

	return nil
}

// Get retrieves a value by key.
func (db *DB) Get(key []byte) ([]byte, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if db.closed {
		return nil, fmt.Errorf("database is closed")
	}

	if val, ok := db.memtable.Get(key); ok {
		return val, nil
	}

	// TODO: search SSTables
	return nil, fmt.Errorf("key not found")
}

// Close shuts down the database.
func (db *DB) Close() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.closed {
		return nil
	}

	db.closed = true
	return db.wal.Close()
}

func (db *DB) flush() error {
	// TODO: write memtable to SSTable, rotate WAL
	db.memtable = newMemTable(db.cfg.MemTableSize)
	return nil
}
