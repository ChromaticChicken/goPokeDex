package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	entries map[string]cacheEntry
	lock    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) PokeCache {
	pC := PokeCache{
		entries: map[string]cacheEntry{},
		lock:    &sync.RWMutex{},
	}
	go pC.reapLoop(interval)
	return pC
}

func (pC PokeCache) Add(key string, value []byte) {
	cE := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	pC.lock.Lock()
	pC.entries[key] = cE
	pC.lock.Unlock()
}

func (pC PokeCache) Get(key string) ([]byte, bool) {
	pC.lock.RLock()
	defer pC.lock.RUnlock()
	if val, ok := pC.entries[key]; ok {
		return val.val, ok
	} else {
		return make([]byte, 0), ok
	}
}

func (pC PokeCache) reapLoop(interval time.Duration) {
	timer := time.NewTicker(interval)
	for reapTime := range timer.C {
		for key, entry := range pC.entries {
			if entry.createdAt.Before(reapTime) {
				pC.lock.Lock()
				delete(pC.entries, key)
				pC.lock.Unlock()
			}
		}
	}
}
