package main

import (
	"fmt"
)

//foo - call by reference
func foo(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

func main() {
	a := [3]int{1, 2, 3}
	foo(&a)
	fmt.Print(a)
}
