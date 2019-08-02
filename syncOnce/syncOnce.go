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

func dostuff(wg *sync.WaitGroup) {
	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}

func main() {
	wg.Add(2)
	go dostuff(&wg)
	go dostuff(&wg)
	wg.Wait()
}
