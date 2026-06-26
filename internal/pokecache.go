package internal

import (
	"time"
	"sync"
)

type Cache struct{
	road map[string]cacheEntry;
	mutex sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time;
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache {
		road: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.road[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	} 

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	for mkey := range c.road {
		if mkey == key {
			return c.road[key].val, true
		}
	}
	
	var result []byte
	return result, false
} 

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {

		c.mutex.Lock()

		for key, entry := range c.road {

			if time.Since(entry.createdAt) > interval {
				delete(c.road, key)
			}
		}

		c.mutex.Unlock()
	}
}