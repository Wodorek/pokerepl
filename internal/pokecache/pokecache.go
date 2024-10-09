package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	cache := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache

}

func (c *Cache) Add(key string, val []byte) {

	entry := cacheEntry{createdAt: time.Now(), val: val}
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = entry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	val, ok := c.cache[key]

	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-last)) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) PrintCache() {
	items := c.cache

	for item := range items {
		fmt.Println(item)
	}
}
