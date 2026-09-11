package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         *sync.RWMutex
	interval   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{interval: interval, cacheEntry: make(map[string]cacheEntry), mu: &sync.RWMutex{}}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock() //when writing we lock
	defer c.mu.Unlock()
	c.cacheEntry[key] = cacheEntry{val: val, createdAt: time.Now()}

}

func (c *Cache) Get(key string) (value []byte, exists bool) {
	c.mu.RLock() //when reading we Rlock
	defer c.mu.RUnlock()
	entry, exists := c.cacheEntry[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() { // Goes through the cache's entries and removes entries older t
	// go routine function in newcache - so it clears cacheentries constantly after being made
	ticker := time.NewTicker(c.interval)

	for range ticker.C { // every 5 seconds, we do perform
		c.mu.Lock()
		// here we loop thru the map and delete stale entries
		for key, entry := range c.cacheEntry {
			if time.Since(entry.createdAt) > c.interval { // if lifespan > interval, remove that entry
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
