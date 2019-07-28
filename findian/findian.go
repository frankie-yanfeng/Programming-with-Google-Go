package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please, enter a string: ")

	var e string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	e = scanner.Text()

	//if err != nil {
	//fmt.Println("Error: ", err)
	//} else {
	fmt.Printf("Input is %s\n", e)
	a := strings.ToLower(e)
	b := strings.TrimSpace(a)
	if strings.HasPrefix(b, "i") && strings.Contains(b[1:len(b)-1], "a") && strings.HasSuffix(b, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
	//}
}
