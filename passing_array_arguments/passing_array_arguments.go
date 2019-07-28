package main

import (
	"fmt"
)

//foo - call by reference
func foo(x [3]int) int {
	return x[0]
}

func main() {
	a := [3]int{1, 2, 3}
	fmt.Print(foo(a))
}
