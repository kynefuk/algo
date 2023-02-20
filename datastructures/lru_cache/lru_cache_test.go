package main

import (
	"fmt"
	"testing"
)

func Test_PutAndGet(t *testing.T) {
	capacity := 3
	cacheKey := 0
	expectedValue := "test"
	cache := NewLRUCache(capacity)
	cache.Put(cacheKey, expectedValue)

	item, err := cache.Get(cacheKey)
	if err != nil {
		t.Fatal(err)
	}

	value, ok := item.(string)
	if !ok {
		t.Fatal(err)
	}

	if value != expectedValue {
		t.Fatal(err)
	}
}

func Test_CapacityLimit(t *testing.T) {
	capacity := 10
	expectedErr := "no item corresponding specified key"
	cache := NewLRUCache(capacity)
	for i := 0; i < capacity; i++ {
		cache.Put(i, struct{}{})
	}

	if _, err := cache.Get(capacity + 1); err.Error() != expectedErr {
		t.Fatalf("expected %s, but got: %s", expectedErr, err)
	}
}

func Test_GetEvicted(t *testing.T) {
	capacity := 2
	expectedErr := "no item corresponding specified key"
	cache := NewLRUCache(capacity)
	for i := 0; i < capacity; i++ {
		str := fmt.Sprintf("test item %d", i)
		cache.Put(i, str)
	}

	_, err := cache.Get(0)
	if err != nil {
		t.Fatal(err)
	}

	cache.Put(capacity+1, fmt.Sprintf("test item %d", capacity+1))
	if _, err = cache.Get(1); err.Error() != expectedErr {
		t.Fatalf("expected %s, but got: %s", expectedErr, err)
	}
}
