package lru1

import "testing"

func TestCache(t *testing.T) {
	size := 4
	cache := NewCache(size)
	items := []*Item{
		&Item{"A", "a"},
		&Item{"B", "b"},
		&Item{"C", "c"},
		&Item{"D", "d"},
		&Item{"E", "e"},
		&Item{"F", "f"},
	}
	for _, i := range items {
		cache.Put(i.key, i.value)
	}

	if cache.lst.Len() != size {
		t.Errorf("got %d list size but expected %d", cache.lst.Len(), size)
	}
	if len(cache.data) != size {
		t.Errorf("got %d map size but expected %d", len(cache.data), size)
	}

	for _, i := range items[2:] {
		c := cache.Get(i.key)
		if *c != *i {
			t.Errorf("got %v item but expected %v", *c, *i)
		}
	}
	for _, i := range items[:2] {
		if c := cache.Get(i.key); c != nil {
			t.Errorf("got %v item but expected %v", *c, nil)
		}
	}
	// TODO: add check order of items in cache
}
