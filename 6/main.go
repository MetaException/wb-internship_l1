package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func first(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение через контекст ")
			return
		default:
			fmt.Println("Работает первая горутина")
			time.Sleep(time.Second * 1)
		}
	}
}

func second(endch chan struct{}) {
	for {
		select {
		case <-endch:
			fmt.Println("Заверешение через канал закрытия")
			return
		default:
			fmt.Println("Работает вторая горутина")
			time.Sleep(time.Second * 1)
		}
	}
}

func third(wg *sync.WaitGroup) {
	fmt.Println("Работает третья горутина")
	time.Sleep(time.Second * 5)
	fmt.Println("Обычное завершение")
	wg.Done()
	// return - обычное завершение
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go first(ctx)

	cancelChan := make(chan struct{})
	go second(cancelChan)

	var wg sync.WaitGroup
	wg.Add(1)

	go third(&wg)

	time.Sleep(time.Second * 3)
	cancel()
	close(cancelChan)
	wg.Wait()
}
