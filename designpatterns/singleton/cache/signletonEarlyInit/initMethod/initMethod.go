package main

// Inspired from - https://medium.com/@ishagirdhar/singleton-pattern-in-golang-9f60d7fdab23

import (
	"fmt"
	"strconv"
	"sync"
)

type cache struct {
	data  map[string]int
	mutex *sync.Mutex
}

var Cache *cache
var mutex *sync.Mutex

func (c *cache) insert(key string, value int) {
	c.mutex.Lock()
	c.data[key] = value
	c.mutex.Unlock()
}

func (c *cache) get(key string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	res, ok := c.data[key]
	if ok {
		return res
	}
	return -1
}

func (c *cache) print() {
	c.mutex.Lock()
	fmt.Println(c)
	c.mutex.Unlock()
}

func init() {
	mutex = &sync.Mutex{}
	Cache = &cache{
		data:  map[string]int{},
		mutex: mutex,
	}
}

func GetInstance() *cache {
	return Cache
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			cacheinstanec := GetInstance()
			cacheinstanec.insert("key"+strconv.Itoa(i), i)
			cacheinstanec.print()
			wg.Done()
		}(i)
	}
	wg.Wait()
	newcacheinstance := GetInstance()
	fmt.Println(newcacheinstance.get("key2"))
}
