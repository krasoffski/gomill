// Package lru1 implement Least Recently Used based on doubly linked list.
package lru1

import "container/list"

// Item is an element in cache.
type Item struct {
	key   string
	value string
}

// Cache is a sized LRU cache.
type Cache struct {
	size int
	data map[string]*list.Element
	lst  *list.List
}

// NewCache returns an initialized LRU cache.
func NewCache(size int) *Cache {
	cache := new(Cache)
	cache.size = size
	cache.data = make(map[string]*list.Element)
	cache.lst = list.New()
	return cache
}

// Put inserts new Item to cache.
// If cache is full removes oldest Item first.
func (c *Cache) Put(key, value string) {
	if len(c.data) == c.size {
		delete(c.data, c.lst.Back().Value.(*Item).key)
		c.lst.Remove(c.lst.Back())
	}
	c.data[key] = c.lst.PushFront(&Item{key: key, value: value})
}

// Get returns Item from cache by key.
// nil is returned if there is no such key in the cache.
func (c *Cache) Get(key string) *Item {
	if c.data[key] != nil {
		c.lst.MoveToFront(c.data[key])
		return c.data[key].Value.(*Item)
	}
	return nil
}
