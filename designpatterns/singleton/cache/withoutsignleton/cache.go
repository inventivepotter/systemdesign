package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Cache struct {
	cache map[string]int
}

type Cacher interface {
	insert(key string, value int)
	get(key string) int
}

func (c *Cache) insert(key string, value int) {
	c.cache[key] = value

}

func (c *Cache) get(key string) int {
	res, ok := c.cache[key]
	if ok {
		return res
	}
	return -1
}

func main() {
	var c Cacher
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			// this will result in concurrent map writes
			// also does miss few elements in the map
			if c == nil {
				c = &Cache{
					cache: map[string]int{},
				}
			}
			c.insert("key"+strconv.Itoa(i), i)
			fmt.Println(i, c)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(c.get("k"))
}
