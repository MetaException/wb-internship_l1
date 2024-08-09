package main

import "fmt"

func main() {
	slice := []int{1, 5, 9, 8, 10}

	var i int
	fmt.Scanln(&i)

	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println(slice)
}
