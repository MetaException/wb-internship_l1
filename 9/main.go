package main

import (
	"fmt"
	"sync"
)

func main() {
	inCh := make(chan int)
	outCh := make(chan int)

	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(2)
	go func(inCh chan int, outCh chan int) {
		defer wg.Done()

		for v := range inCh {
			outCh <- v * v
		}

		close(outCh)
	}(inCh, outCh)

	go func(outCh chan int) {
		defer wg.Done()
		for v := range outCh {
			fmt.Println(v)
		}
	}(outCh)

	for _, v := range arr {
		inCh <- v
	}
	close(inCh)

	wg.Wait()
}
