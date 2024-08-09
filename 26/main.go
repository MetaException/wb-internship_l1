package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Scanln(&input)

	fmt.Println(CheckUnique(input))
}

func CheckUnique(str string) bool {
	hashset := make(map[rune]struct{})

	for _, v := range str {
		if _, ok := hashset[v]; ok {
			return false
		}
		hashset[unicode.ToLower(v)] = struct{}{}
	}
	return true
}
