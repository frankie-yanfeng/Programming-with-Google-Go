package main

import (
	"fmt"
	"sync"
)

var i int
var wg sync.WaitGroup

func inc() {
	i = i + 1
	wg.Done()
}

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}
