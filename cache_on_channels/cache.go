package cache

import (
	"sync"
)

type cacher interface {
	set(key int, val interface{})
	read(key int) interface{}
}

// Channels based cache.
//
//
type payload struct {
	status int // 0 - reader; 1 - setter
	key    int
	val    interface{}
}

type cache struct {
	items map[int]interface{}

	in   chan payload
	outS chan struct{}
	outR chan interface{}
}

func newCache() cacher {
	c := &cache{
		items: make(map[int]interface{}),

		in:   make(chan payload),
		outS: make(chan struct{}),
		outR: make(chan interface{}),
	}
	go handler(c)
	return c
}

func handler(c *cache) {
	for p := range c.in {
		switch p.status {
		case 0:
			if val, ok := c.items[p.key]; ok {
				c.outR <- val
			} else {
				c.outR <- nil
			}
		case 1:
			c.items[p.key] = p.val
			// c.outS <- struct{}{}
		}
	}
}

func (c *cache) set(key int, val interface{}) {
	c.in <- payload{1, key, val}
	// <-c.outS
}

func (c *cache) read(key int) interface{} {
	c.in <- payload{status: 0, key: key}
	return <-c.outR
}

// Mutex based cache.
//
//
type cacheM struct {
	items map[int]interface{}
	sync.RWMutex
}

func newCacheM() cacher {
	return &cacheM{items: make(map[int]interface{})}
}

func (c *cacheM) set(key int, val interface{}) {
	c.Lock()
	c.items[key] = val
	c.Unlock()
}

func (c *cacheM) read(key int) interface{} {
	c.RLock()
	if val, ok := c.items[key]; ok {
		c.RUnlock()
		return val
	}
	c.RUnlock()
	return nil
}
