package main

import (
	"fmt"
	"sync"
)

func main() {
	var ans int
	nums := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			ans += n * n
			mu.Unlock()
		}(num)
	}

	wg.Wait()

	fmt.Println(ans)
}
