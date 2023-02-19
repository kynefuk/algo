package main

import (
	"fmt"
	"math"
	"sync"
)

type item struct {
	value interface{}
	age   int
}

type LRUCache struct {
	capacity   int
	currentAge int
	items      map[int]*item
	lock       *sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:   capacity,
		currentAge: 0,
		items:      make(map[int]*item, capacity),
		lock:       new(sync.Mutex),
	}
}

func (c *LRUCache) Get(key int) (interface{}, error) {
	if item, ok := c.items[key]; ok {
		c.lock.Lock()
		defer c.lock.Unlock()
		item.age = c.currentAge
		c.currentAge++
		return item.value, nil
	}
	return nil, fmt.Errorf("no item corresponding specified key")
}

func (c *LRUCache) Put(key int, value interface{}) error {
	if item, ok := c.items[key]; ok {
		c.lock.Lock()
		defer c.lock.Unlock()
		item.value = value
		item.age = c.currentAge
		c.currentAge++
		return nil
	}

	if len(c.items) >= c.capacity {
		leastAge := math.MaxInt64
		leastAgeKey := 0
		for key, item := range c.items {
			if item.age < leastAge {
				leastAge = item.age
				leastAgeKey = key
			}
		}

		if leastAgeKey != 0 {
			delete(c.items, leastAgeKey)
		}
	}
	c.items[key] = &item{
		age:   c.currentAge,
		value: value,
	}
	c.currentAge++

	return nil
}

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, "first cache item")
	cache.Put(2, "second cache item")

	first_cache, err := cache.Get(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(first_cache)

	cache.Put(3, "third cache item") // This operation evicts key `2`

	second_cache, err := cache.Get(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(second_cache)
}
