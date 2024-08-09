package point

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p Point) CalculateDinstance(otherPoint *Point) float64 {
	xDist := otherPoint.x - p.x
	yDist := otherPoint.y - p.y

	return math.Sqrt(xDist*xDist + yDist*yDist)
}
