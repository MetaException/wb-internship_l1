package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		sqr(arr[:(len(arr) / 2)])
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sqr(arr[(len(arr) / 2):])
	}()

	wg.Wait()
	fmt.Println(arr)
}

func sqr(arr []int) []int {
	for i, v := range arr {
		arr[i] = v * v
	}
	return arr
}
