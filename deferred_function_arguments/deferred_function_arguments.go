package main

import (
	"fmt"
)

func main() {
	var i = 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println("Hello!")
}
