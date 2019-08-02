package main

import (
	"fmt"
	"sync"
)

var i int

//var mut = &sync.Mutex{}

func inc(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	i = i + 1
	m.Unlock()
	wg.Done()
}

func main() {

	var w sync.WaitGroup
	var m sync.Mutex

	w.Add(2)
	go inc(&w, &m)
	go inc(&w, &m)
	w.Wait()
	fmt.Printf("i is %d\n", i)
}
