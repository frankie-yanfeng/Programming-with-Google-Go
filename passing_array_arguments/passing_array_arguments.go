package main

import (
	"fmt"
)

//foo - call by value
func foo(x [3]int) {
	x[0] = x[0] + 1

	fmt.Println("modified array is: ", x)
}

func main() {
	a := [3]int{1, 2, 3}
	foo(a)

	fmt.Println("original array is: ", a)
}
