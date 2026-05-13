package wal

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

// Writer handles write-ahead log operations.
type Writer struct {
	file *os.File
	sync bool
}

// NewWriter creates a new WAL writer.
func NewWriter(dataDir string, sync bool) (*Writer, error) {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("create data dir: %w", err)
	}

	path := filepath.Join(dataDir, "wal.log")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("open wal file: %w", err)
	}

	return &Writer{file: f, sync: sync}, nil
}

// Append writes a key-value pair to the WAL.
func (w *Writer) Append(key, value []byte) error {
	// Format: [key_len:4][value_len:4][key][value]
	header := make([]byte, 8)
	binary.LittleEndian.PutUint32(header[0:4], uint32(len(key)))
	binary.LittleEndian.PutUint32(header[4:8], uint32(len(value)))

	if _, err := w.file.Write(header); err != nil {
		return err
	}
	if _, err := w.file.Write(key); err != nil {
		return err
	}
	if _, err := w.file.Write(value); err != nil {
		return err
	}

	if w.sync {
		return w.file.Sync()
	}
	return nil
}

// Close closes the WAL file.
func (w *Writer) Close() error {
	if w.file != nil {
		return w.file.Close()
	}
	return nil
}
