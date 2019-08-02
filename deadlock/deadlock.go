package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var on sync.Once

func setup() {
	fmt.Println("Init")
}

func dostuff(c1 chan int, c2 chan int) {
	<-c1
	c2 <- 1
	wg.Done()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2)
	go dostuff(ch1, ch2)
	go dostuff(ch2, ch1)
	wg.Wait()
}
