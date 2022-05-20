package cache

import "time"

type Value struct {
	value    string
	deadline time.Time
}

type Cache struct {
	Values map[string]Value
}

func NewCache() Cache {
	return Cache{map[string]Value{}}
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.Values[key]
	if !ok {
		return "", false
	}

	if !value.deadline.IsZero() && value.deadline.Before(time.Now()) {
		delete(c.Values, key)
		return "", false
	}
	return value.value, true
}

func (c *Cache) Put(key, value string) {
	var t time.Time
	c.Values[key] = Value{value, t}
}

func (c *Cache) Keys() []string {
	keys := []string{}
	for k, v := range c.Values {
		if v.deadline.IsZero() || v.deadline.Before(time.Now()) {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	if deadline.After(time.Now()) {
		c.Values[key] = Value{value, deadline}
	}
}
