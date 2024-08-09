package main

import (
	"fmt"
)

func quicksort(arr []int, left int, right int) {
	if left < right {
		p := partition(arr, left, right)
		quicksort(arr, left, p-1)
		quicksort(arr, p+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	p := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < p {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main() {
	var arr = []int{5, 9, 6, 2, 0, 1, 4, 5, 8, 9}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
