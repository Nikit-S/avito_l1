package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	val uint64
	sync.RWMutex
}

func (c *Counter) Get() uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.val
}

func (c *Counter) Increment(i uint64) {
	c.Lock()
	c.val += i
	c.Unlock()
}

func (c *Counter) Set(i uint64) {
	c.Lock()
	c.val = i
	c.Unlock()
}

//такой же принцип что и с мапой в конкурентной среде
func main() {
	var cnt Counter
	cnt.Set(0)

	var wg sync.WaitGroup
	wg.Add(3)

	go func(c *Counter) {
		for i := 0; i < 5; i++ {
			c.Increment((uint64)(i))
		}
		wg.Done()
	}(&cnt)

	go func(c *Counter) {
		for i := 0; i < 5; i++ {
			c.Increment((uint64)(i * 100))
		}
		wg.Done()
	}(&cnt)

	go func(c *Counter) {
		for i := 0; i < 5; i++ {
			c.Increment((uint64)(i * 10))
		}
		wg.Done()
	}(&cnt)
	wg.Wait()
	fmt.Println(cnt.Get())

}
