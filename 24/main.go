package main

import (
	"fmt"
	"point/point"
)

func main() {
	point1 := point.NewPoint(0, 0)
	point2 := point.NewPoint(1, 1)

	fmt.Println(point1.CalculateDinstance(point2))
}
