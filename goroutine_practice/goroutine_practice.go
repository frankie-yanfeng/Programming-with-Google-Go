package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

//PrintX definition
func sortArray(x []int, wg *sync.WaitGroup) {
	sort.Ints(x)
	wg.Done()
}

func main() {
	var unit int
	var intArray []int
	intArray = make([]int, 0)

	var intArrayResult []int
	intArrayResult = make([]int, 0)

	for {
		fmt.Printf("\n>please enter an array of integers for sorting using space as the delimiter: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		input = strings.Trim(input, " ")

		fmt.Printf("input is: [%s]\n", input)

		if len(strings.Split(input, " ")) < 4 {
			fmt.Println("integer input number is less than 4, please input equal or more than 4 integers")
			continue
		} else {
			fmt.Printf("The input length is %d\n", len(input))

			for _, v := range strings.Split(input, " ") {
				//fmt.Printf("index =%d , value =%s\n", i, v)
				n, _ := strconv.Atoi(v)

				intArray = append(intArray, n)
			}
			fmt.Println("integer array is:", intArray)
			fmt.Println("integer array len is:", len(intArray))
			break
		}
	}

	unit = len(intArray) / 4

	fmt.Printf("unit is: %d\n", unit)

	iniArrayA := intArray[0:unit]
	iniArrayB := intArray[unit : unit*2]
	iniArrayC := intArray[unit*2 : unit*3]
	iniArrayD := intArray[unit*3:]

	fmt.Println("iniArrayA array is:", iniArrayA)
	fmt.Println("iniArrayB array is:", iniArrayB)
	fmt.Println("iniArrayC array is:", iniArrayC)
	fmt.Println("iniArrayD array is:", iniArrayD)

	var wg sync.WaitGroup
	wg.Add(4)
	go sortArray(iniArrayA, &wg)
	go sortArray(iniArrayB, &wg)
	go sortArray(iniArrayC, &wg)
	go sortArray(iniArrayD, &wg)
	wg.Wait()

	fmt.Println()

	fmt.Println("iniArrayA array is:", iniArrayA)
	fmt.Println("iniArrayB array is:", iniArrayB)
	fmt.Println("iniArrayC array is:", iniArrayC)
	fmt.Println("iniArrayD array is:", iniArrayD)

	intArrayResult = append(intArrayResult, iniArrayA...)
	intArrayResult = append(intArrayResult, iniArrayB...)
	intArrayResult = append(intArrayResult, iniArrayC...)
	intArrayResult = append(intArrayResult, iniArrayD...)

	fmt.Println("intArrayResult before sorting: ", intArrayResult)
	sort.Ints(intArrayResult)
	fmt.Println("intArrayResult after  sorting: ", intArrayResult)
	// fmt.Println(a * b)
}
