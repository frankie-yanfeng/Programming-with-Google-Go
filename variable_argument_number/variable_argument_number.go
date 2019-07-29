package main

import "fmt"

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}

	return maxV
}

func main() {
	fmt.Println(getMax(1, 3, 6, 4))
	vslice := []int{1, 3, 6, 4}
	fmt.Println(getMax(vslice...))
}
