package main

import (
	"fmt"
	"sync"
)

//goo definition
func goo(wg1 *sync.WaitGroup) {
	fmt.Println("inside goo")
	wg1.Done()
}

//foo definition
func foo(wg *sync.WaitGroup) {

	var wg1 sync.WaitGroup

	wg1.Add(1)
	goo(&wg1)
	wg1.Wait()

	fmt.Println("inside foo")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()

	fmt.Println("inside main")
}
