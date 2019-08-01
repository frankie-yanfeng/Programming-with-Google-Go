package main

import "fmt"

//Point definition
type Point struct {
	x float64
	y float64
}

//OffsetX definition
func (p *Point) OffsetX(v float64) {
	fmt.Println(p.x)
	p.x = p.x + v

	fmt.Println(p.x)
}

func main() {
	m := Point{3, 4}

	m.OffsetX(5)

	fmt.Println(m.x)
}
