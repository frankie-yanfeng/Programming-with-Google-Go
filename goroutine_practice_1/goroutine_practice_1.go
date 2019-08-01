package main

import (
	"fmt"
	"sort"
)

func sortIntSlice(i int, iS []int, ch chan []int) {
	fmt.Printf("goroutine #%d with these integers: %v\n", i, iS)
	sort.Ints(iS)
	ch <- iS
}

func main() {
	var c chan []int
	c = make(chan []int)

	a := []int{1, 2, 5, 4, 3}

	go sortIntSlice(0, a, c)

	result := <-c

	fmt.Println("result: ", result)
}
