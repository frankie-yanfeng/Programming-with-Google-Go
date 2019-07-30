package main

import (
	"fmt"
)

//Speaker body
type Speaker interface {
	Speak()
}

//Dog body
type Dog struct{ name string }

//Speak body
func (d Dog) Speak() { fmt.Println(d.name) }

func main() {
	var s1 Speaker
	d1 := Dog{"Brian"}
	s1 = d1
	s1.Speak()
}
