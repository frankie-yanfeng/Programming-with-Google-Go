package main

import (
	"fmt"
)

//Speaker interface
type Speaker interface {
	Speak()
}

//Dog struct
type Dog struct{ name string }

//Speak method
func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

func main() {
	var s1 Speaker
	var d1 *Dog
	s1 = d1
	s1.Speak()
}
