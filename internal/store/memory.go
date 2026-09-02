package store

import "sync"

type ItemStore struct {
	mu    sync.RWMutex
	items map[string]string
}

func NewitemStore() *ItemStore {
	return &ItemStore{items: make(map[string]string)}
}

func (s *ItemStore) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.items[key]
	return val, ok
}

func (s *ItemStore) Set(key, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = val
}
