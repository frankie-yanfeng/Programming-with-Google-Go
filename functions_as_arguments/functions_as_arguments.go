package main

import (
	"fmt"
)

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }

func main() {
	fmt.Print(applyIt(incFn, 2))
	fmt.Print("\n")
	fmt.Print(applyIt(decFn, 2))
}
