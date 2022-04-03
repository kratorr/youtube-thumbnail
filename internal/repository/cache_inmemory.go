package repository

import (
	"sync"
)

type CacheInmemory struct {
	cache map[string][]byte
	mu    *sync.RWMutex
}

func NewCacheInmemory(cache map[string][]byte, mu *sync.RWMutex) *CacheInmemory {
	return &CacheInmemory{cache: cache, mu: mu}
}

func (r *CacheInmemory) InCache(videoID string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, ok := r.cache[videoID]
	if ok {
		return true
	}
	return false
}

func (r *CacheInmemory) Get(videoID string) []byte {
	r.mu.RLock()
	defer r.mu.RUnlock()
	image := r.cache[videoID]

	return image
}

func (r *CacheInmemory) Insert(videoID string, image []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cache[videoID] = image
}
