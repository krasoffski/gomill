package lru1

import "container/list"

type Item struct {
	key   string
	value string
}

type Cache struct {
	capacity int
	data     map[string]*list.Element
	lst      *list.List
}

func NewCache(capacity int) *Cache {
	cache := new(Cache)
	cache.capacity = capacity
	cache.data = make(map[string]*list.Element)
	return cache
}

func (c *Cache) Put(key, value string) {
	if len(c.data) == c.capacity {
		delete(c.data, c.lst.Back().Value.(*Item).key)
		c.lst.Remove(c.lst.Back())
	}
}

func (c *Cache) Get(key string) *Item {
	if c.data[key] != nil {
		c.lst.MoveToFront(c.data[key])
		return c.data[key].Value.(*Item)
	}
	return nil
}
