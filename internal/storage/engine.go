package storage

import "sync"

type Storage interface {
	Set(key string, value string)
	Get(key string) (string, bool)
	Delete(key string)
}

type inMemoryEngine struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryEngine() Storage {
	return &inMemoryEngine{
		data: make(map[string]string),
	}
}

func (e *inMemoryEngine) Set(key string, value string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.data[key] = value
}

func (e *inMemoryEngine) Get(key string) (string, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	val, ok := e.data[key]
	return val, ok
}

func (e *inMemoryEngine) Delete(key string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.data, key)
}
