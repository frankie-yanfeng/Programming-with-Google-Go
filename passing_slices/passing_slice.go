package main

import (
	"fmt"
)

//foo - call by reference
func foo(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	a := []int{1, 2, 3}
	foo(a)
	fmt.Print(a)
}
