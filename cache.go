package cache

import (
	"time"
)

type ValueSt struct {
	value    string
	deadline time.Time
}

type Cache struct {
	store map[string]ValueSt
}

func NewCache() Cache {
	return Cache{
		store: make(map[string]ValueSt),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	val, ok := c.store[key]
	if !ok {
		return "", false
	}

	if !val.deadline.IsZero() {
		if time.Since(val.deadline) >= 0 {
			delete(c.store, key)
			return "", false
		}
	}

	return val.value, true
}

func (c *Cache) Put(key, value string) {
	c.store[key] = ValueSt{
		value: value,
	}
}

func (c *Cache) Keys() []string {
	arr := make([]string, 0, len(c.store))
	for k := range c.store {
		_, ok := c.Get(k)
		if ok {
			arr = append(arr, k)
		}
	}

	return arr
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.store[key] = ValueSt{
		value:    value,
		deadline: deadline,
	}
}
