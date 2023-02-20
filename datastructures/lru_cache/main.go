package main

import (
	"fmt"
)

func main() {
	cache := NewLRUCache(2)
	cache.Put(1, "first cache item")
	cache.Put(2, "second cache item")

	first_cache, err := cache.Get(1)
	if err != nil {
		fmt.Println(err)
	}

	v, ok := first_cache.(string)
	if ok {
		fmt.Println(v)
	}

	cache.Put(3, "third cache item")

	second_cache, err := cache.Get(2)
	if err != nil {
		fmt.Println(err)
	}

	v, ok = second_cache.(string)
	if ok {
		fmt.Println(v)
	}
}
