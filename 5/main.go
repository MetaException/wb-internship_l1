package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var timeout int
	fmt.Scanln(&timeout)

	ch := make(chan int)
	rand := rand.New(rand.NewSource(time.Now().UnixMilli()))

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(timeout)))
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(2)
	go func(c context.Context, inCh chan int) {
		defer wg.Done()
		for {
			select {
			case <-c.Done():
				close(inCh)
				return
			default:
				inCh <- rand.Int()
			}
		}
	}(ctx, ch)

	go func(inCh chan int) {
		defer wg.Done()
		for v := range inCh { // Явл упрощённой формой select + case data, ok := <- ch if !ok return...
			fmt.Println(v)
		}
	}(ch)

	wg.Wait()
}
