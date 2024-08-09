package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	var workersCount int
	fmt.Scan(&workersCount)

	inCh := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go WorkerFunc(&wg, inCh, i)
	}

	rand := rand.New(rand.NewSource(time.Now().UnixMilli()))

	for {
		select {
		case <-done:
			fmt.Println("Выход из программы...")
			close(inCh)
			wg.Wait()
			return

		default:
			inCh <- rand.Int()
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func WorkerFunc(wg *sync.WaitGroup, inCh <-chan int, number int) {
	defer wg.Done()
	for el := range inCh {
		fmt.Println(el, number)
	}
}
