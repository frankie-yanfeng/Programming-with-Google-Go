package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Name struct
type Name struct {
	fname string
	lname string
}

func main() {
	var s []Name
	s = make([]Name, 0)
	var record Name

	fmt.Println("please enter the file name:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	fmt.Printf("filename is: %s\n", filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	con := string(content)

	//fmt.Printf("file content is:\n %s\n", con)

	//fmt.Printf("file line: %s", strings.Split(con, "\n")[0])

	for i, l := range strings.Split(con, "\n") {
		fmt.Printf("no: %d, file line: %s\n", i, l)
		record.fname = strings.Split(l, " ")[0]
		record.lname = strings.Split(l, " ")[1]
		s = append(s, record)
	}

	for m, n := range s {
		fmt.Printf("\nline: %d, first name: %s, last name: %s", m, n.fname, n.lname)
	}

	fmt.Printf("\n")
}
