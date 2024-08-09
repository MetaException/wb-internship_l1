package main

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

type ConcurrentMap struct {
	vmap map[string]interface{}
	mu   *sync.Mutex
}

func NewMap() *ConcurrentMap {
	return &ConcurrentMap{
		vmap: make(map[string]interface{}),
		mu:   new(sync.Mutex),
	}
}

func (cmap *ConcurrentMap) Set(key string, value interface{}) {
	cmap.mu.Lock()
	cmap.vmap[key] = value
	cmap.mu.Unlock()
}

func main() {

	cmap := NewMap()

	rand := rand.New(rand.NewSource(time.Now().UnixMilli()))

	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	wg.Add(2)
	go func(ctx context.Context) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				cmap.Set(string(rand.Int()), rand.Int())
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				cmap.Set(string(rand.Int()), rand.Int())
			}
		}
	}(ctx)

	wg.Wait()
}
