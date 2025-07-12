package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu 		*sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val 	  []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache {
		entries: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (ch *Cache) Add(key string, value []byte) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	ch.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (ch *Cache) Get(key string) ([]byte, bool) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	val, ok := ch.entries[key]
	return val.val, ok
}

func (ch *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		ch.cleanUp(time.Now().UTC(), interval)
	}
}

func (ch *Cache) cleanUp(now time.Time, last time.Duration) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	for key, val := range ch.entries {
		if val.createdAt.Before(now.Add(-last)) {
			delete(ch.entries, key)
		}
	}
}
