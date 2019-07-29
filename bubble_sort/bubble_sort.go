package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Swap body
func Swap(sli []int, index int) {
	var tmp int
	tmp = sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = tmp
}

//bubble_sort body
func bubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {
	//a := []int{2, 1, 0, 3, 9, 4, 11, 6, 7, 5}
	var a []int
	var input string

	a = make([]int, 0, 10)

	for i := 0; i < cap(a); i++ {
		fmt.Printf("Please, enter %d input: ", i)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		input = scanner.Text()

		fmt.Printf("%d Input is %s\n", i, input)

		num, _ := strconv.Atoi(input)
		a = append(a, num)
	}

	fmt.Println(a)
	bubbleSort(a)
	fmt.Println(a)
}
