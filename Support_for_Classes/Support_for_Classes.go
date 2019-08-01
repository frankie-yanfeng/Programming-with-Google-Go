package main

import (
	"fmt"
)

//MyInt type
type MyInt int

//Double association
func (mi MyInt) Double() int {
	return int(mi * 2)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())
}
