package kvstore

import (
	"sync"
)

type MemoryKVStore struct {
	mutex sync.RWMutex
	m     map[string]interface{}
}

func (s MemoryKVStore) Get(key string) interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.m[key]
}

func (s MemoryKVStore) Set(key string, value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[key] = value
}

func (s MemoryKVStore) Delete(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.m, key)
}

func (s MemoryKVStore) Disactive() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
}

func (s MemoryKVStore) Active() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
}
