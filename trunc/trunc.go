package main

import (
	"fmt"
)

func main() {
	fmt.Print("Please, enter a float number: ")

	var e float32

	_, err :=fmt.Scanf("%f", &e)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("%f converted to int is %d\n", e, int(e))
	}
}
