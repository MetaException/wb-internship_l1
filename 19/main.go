package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanln(&input)

	result := []rune(input)

	for i := len(input) - 1; i >= 0; i-- {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	fmt.Println(string(result))
}
