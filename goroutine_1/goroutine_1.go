package main

import (
	"fmt"
	"time"
)

//PrintX definition
func PrintX() {
	fmt.Println("inside goroutine")
}

func main() {
	go PrintX()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("inside main routine")
}
