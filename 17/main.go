package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 6, 9, 20}

	fmt.Println(binarySearch(arr[:], 20))
}

func binarySearch(arr []int, value int) int {

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if value > arr[mid] {
			left = mid + 1
		} else if value < arr[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
