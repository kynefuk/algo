package main

type LRUCacheIF interface {
	Get(key int) (interface{}, error)
	Put(key int, value interface{}) error
}
