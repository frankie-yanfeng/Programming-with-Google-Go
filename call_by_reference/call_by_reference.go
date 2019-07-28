package main

import (
	"fmt"
)

//foa - call by value
func foa(y int) {
	y = y + 1
	fmt.Printf("value:%d\n", y)
}

//foo - call by reference
func foo(y *int) {
	*y = *y + 1
}

func main() {
	x := 2
	foa(x)
	fmt.Printf("value:%d\n", x)

	foo(&x)
	fmt.Printf("value:%d\n", x)
}
