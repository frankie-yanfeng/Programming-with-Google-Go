package main

import (
	"fmt"
	"math"
)

//Point definition
type Point struct {
	x float64
	y float64
}

//DistToOrig body
func (p Point) DistToOrig() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)

	return math.Sqrt(t)
}

func main() {
	p1 := Point{3, 4}
	fmt.Println(p1.DistToOrig())
}
