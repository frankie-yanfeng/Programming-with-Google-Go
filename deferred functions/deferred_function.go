package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Bye")

	fmt.Println("Hello!")
}
