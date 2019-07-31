package main

import (
	"fmt"
	"time"
)

var count int

func race() {
	fmt.Println(count)
	count++
}

func main() {

	count = 1
	go race()
	go race()
	time.Sleep(1 * time.Second)
}
