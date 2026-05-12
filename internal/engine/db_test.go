package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDB_PutGet(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{
		DataDir:      dir,
		WALSync:      false,
		MemTableSize: 1024,
		BloomFPRate:  0.01,
	}

	db, err := Open(cfg)
	require.NoError(t, err)
	defer db.Close()

	err = db.Put([]byte("hello"), []byte("world"))
	require.NoError(t, err)

	val, err := db.Get([]byte("hello"))
	require.NoError(t, err)
	assert.Equal(t, []byte("world"), val)
}

func TestDB_GetMissing(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{
		DataDir:      dir,
		WALSync:      false,
		MemTableSize: 1024,
		BloomFPRate:  0.01,
	}

	db, err := Open(cfg)
	require.NoError(t, err)
	defer db.Close()

	_, err = db.Get([]byte("missing"))
	assert.Error(t, err)
}

func BenchmarkDB_Put(b *testing.B) {
	dir := b.TempDir()
	cfg := Config{
		DataDir:      dir,
		WALSync:      false,
		MemTableSize: 64 * 1024 * 1024,
		BloomFPRate:  0.01,
	}

	db, err := Open(cfg)
	require.NoError(b, err)
	defer db.Close()

	key := []byte("bench-key")
	val := make([]byte, 256)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.Put(key, val)
	}
}
