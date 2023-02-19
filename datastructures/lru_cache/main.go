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
	fmt.Println(first_cache)

	cache.Put(3, "third cache item") // This operation evicts key `2`

	second_cache, err := cache.Get(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(second_cache)
}
