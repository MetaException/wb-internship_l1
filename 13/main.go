package main

import "fmt"

func main() {

	/*
		nums := []int{1, 2}

		fmt.Println(nums)

		nums[0], nums[1] = nums[1], nums[0] // Всё равно внутри будет создана временные переменные

		fmt.Println(nums)
	*/

	a := 99
	b := 55

	fmt.Println(a, b)

	a += b
	b = a - b
	a -= b

	fmt.Println(a, b)
}
