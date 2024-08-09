package main

import "fmt"

func main() {
	A := make(map[int]struct{})
	B := make(map[int]struct{})
	C := make(map[int]struct{})

	for i := 0; i <= 10; i++ {
		A[i*2] = struct{}{}
		B[i+5] = struct{}{}
	}

	for k := range A {
		if _, ok := B[k]; ok {
			C[k] = struct{}{}
		}
	}

	fmt.Println(C)
}
