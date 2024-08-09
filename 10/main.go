package main

import (
	"fmt"
)

func main() {
	nums := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float32)

	for _, v := range nums {
		group := int(v/10) * 10
		groups[group] = append(groups[group], v)
	}

	fmt.Println(groups)
}
