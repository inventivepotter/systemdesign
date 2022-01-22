package main

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

func GetInstance(m *sync.Mutex) *cache {
	// Double Lock Signleton Implementation
	if Cache == nil {
		m.Lock()
		if Cache == nil {
			Cache = &cache{
				data:  map[string]int{},
				mutex: m,
			}
		}
		m.Unlock()
	}
	return Cache
}

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			cacheinstance := GetInstance(&mutex)
			cacheinstance.insert("key"+strconv.Itoa(i), i)
			cacheinstance.print()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
