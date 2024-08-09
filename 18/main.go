package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value atomic.Int32
}

func NewCounter() *Counter {
	return &Counter{
		value: atomic.Int32{},
	}
}

func (c *Counter) Increment() {
	c.value.Add(1)
}

func (c *Counter) Get() int32 {
	return c.value.Load()
}

func main() {
	var gorutinesCount int
	fmt.Scan(&gorutinesCount)

	var wg sync.WaitGroup

	c := NewCounter()
	for i := 0; i < gorutinesCount; i++ {
		wg.Add(1)
		go func() {
			c.Increment()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(c.Get())
}
