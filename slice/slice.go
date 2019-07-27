package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	var s []int

	s = make([]int, 0, 3)

	fmt.Printf("s len:%d, s cap:%d, s: %v\n", len(s), cap(s), s)

	for {
		fmt.Print("Please, enter an input: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		input = scanner.Text()

		fmt.Printf("Input is %s\n", input)

		if strings.ToUpper(input)[0] == 'X' {
			fmt.Printf("X is inputed, over\n")
			fmt.Printf("sorted slice %v\n", s)
			fmt.Printf("s len:%d, s cap:%d, s: %v\n", len(s), cap(s), s)
			return
		}

		//num, _ := strconv.ParseInt(input, 0, 32)
		num, _ := strconv.Atoi(input)

		s = append(s, num)
		sort.Ints(s)
		//fmt.Printf("sorted slice %v\n", s)
	}
}
