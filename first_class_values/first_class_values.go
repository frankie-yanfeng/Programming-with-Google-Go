package main

import (
	"fmt"
)

var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func main() {
	funcVar = incFn
	fmt.Print(funcVar(1))
}
