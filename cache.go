package cache

import "time"

type ValueSt struct {
	value       string
	deadline    time.Time
	hasDeadline bool
}

type Cache struct {
	store map[string]ValueSt
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	val, ok := c.store[key]
	if !ok {
		return "", false
	}

	if val.hasDeadline {
		if time.Since(val.deadline) < 0 {
			return "", false
		}
	}

	return val.value, true
}

func (c *Cache) Put(key, value string) {
	c.store[key] = ValueSt{value: value}
}

func (c *Cache) Keys() []string {
	arr := make([]string, 0, len(c.store))
	for k := range c.store {
		val, ok := c.Get(k)
		if ok {
			arr = append(arr, val)
		} else {
			delete(c.store, k)
		}
	}

	return arr
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.store[key] = ValueSt{
		value:       value,
		deadline:    deadline,
		hasDeadline: true,
	}
}
