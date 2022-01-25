package main

// Inspired by - https://www.saylorsolutions.com/singleton-pattern-in-golang/ and
//Inspired by - https://www.saylorsolutions.com/singleton-pattern-in-golang/

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
var mutex = sync.Mutex{}

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

func GetInstance() *cache {
	// Double Lock Signleton Implementation
	if Cache == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if Cache == nil {
			Cache = &cache{
				data:  map[string]int{},
				mutex: &mutex,
			}
			fmt.Println("cache created")
		} else {
			fmt.Println("Cache used")
		}
	} else {
		fmt.Println("Cache used")
	}
	return Cache
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			cacheinstance := GetInstance()
			cacheinstance.insert("key"+strconv.Itoa(i), i)
			cacheinstance.print()
			wg.Done()
		}(i)
	}
	wg.Wait()
	cacheinstance := GetInstance()
	fmt.Println(cacheinstance.get("key2"))
}
