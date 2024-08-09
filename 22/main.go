package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, result big.Int
	fmt.Scan(&a, &b)

	fmt.Println(result.Add(&a, &b))
	fmt.Println(result.Sub(&a, &b))
	fmt.Println(result.Mul(&a, &b))
	fmt.Println(result.Div(&a, &b))
}
