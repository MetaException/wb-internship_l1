package main

import "fmt"

func main() {
	var num, i int64
	fmt.Scan(&num, &i)

	var bitmask int64 = 1 << i

	fmt.Printf("%064b\n", num^bitmask)
}
