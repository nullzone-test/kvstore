package engine

import "sync"

type memTable struct {
	mu      sync.RWMutex
	data    map[string][]byte
	size    int
	maxSize int
}

func newMemTable(maxSize int) *memTable {
	return &memTable{
		data:    make(map[string][]byte),
		maxSize: maxSize,
	}
}

func (m *memTable) Put(key, value []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()

	prev, exists := m.data[string(key)]
	m.data[string(key)] = value

	if exists {
		m.size -= len(prev)
	}
	m.size += len(key) + len(value)
}

func (m *memTable) Get(key []byte) ([]byte, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.data[string(key)]
	return val, ok
}

func (m *memTable) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.size
}
