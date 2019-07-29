package main

import (
	"fmt"
	"math"
)

//GenDisplaceFn body
func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(time float64) float64 {
	fn := func(t float64) float64 {
		return 1/2*acceleration*math.Pow(t, 2) + initialVelocity*t + initialDisplacement
	}
	return fn
}

func main() {
	fn := GenDisplaceFn(10, 2, 1)
	fmt.Println(fn(5))
}
