package main

import "fmt"

func main() {
	var toDectect interface{}

	toDectect = "fwew"

	if _, ok := toDectect.(int); ok {
		fmt.Println("Это int")
	} else if _, ok := toDectect.(string); ok {
		fmt.Println("Это string")
	} else if _, ok := toDectect.(bool); ok {
		fmt.Println("Это bool")
	} else if _, ok := toDectect.(chan interface{}); ok {
		fmt.Println("Это канал")
	} else {
		fmt.Println("Неизвестный тип")
	}
}
