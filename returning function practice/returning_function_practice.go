package main

import (
	"fmt"
	"math"
	"strconv"
)

//GenDisplaceFn body
func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(time float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*acceleration*math.Pow(t, 2) + initialVelocity*t + initialDisplacement
	}
	return fn
}

func main() {
	var acceleration, initialVelocity, initialDisplacement, time float64

	fmt.Println("Welcome. This program will compute the displacement after an specific time.")

	fmt.Printf("Please enter the value for the acceleration: ")
	_, err := fmt.Scan(&acceleration)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Please enter the value for the initial velocity: ")
	_, err = fmt.Scan(&initialVelocity)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Please enter the value for the initial displacement: ")
	_, err = fmt.Scan(&initialDisplacement)
	if err != nil {
		fmt.Println(err)
	}

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	for {
		fmt.Printf("\nPlease enter a time value to compute the displacement: ")

		_, err := fmt.Scan(&time)

		if err != nil {
			fmt.Println(err)
		}

		displacementComputed := fn(time)

		fmt.Printf("-- Displacement computed: %s\n\n", strconv.FormatFloat(displacementComputed, 'f', 2, 64))
	}
}
