package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal definition
type Animal struct {
	food, locomotion, noise string
}

//Eat definition
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

//Move definition
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

//Speak definition
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {

	var animal Animal

	for {
		fmt.Printf(">please enter the request -> name & infomation: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()

		//fmt.Printf("request is: %s\n", request)

		name := strings.Split(request, " ")[0]
		info := strings.Split(request, " ")[1]

		//fmt.Printf("name: %s\n", name)
		//fmt.Printf("info: %s\n", info)

		switch name {
		case "cow":
			animal = Animal{"grass", "walk", "moo"}
		case "bird":
			animal = Animal{"worms", "fly", "peep"}
		case "snake":
			animal = Animal{"mice", "slither", "hsss"}
		default:
			fmt.Printf("%s is not in (cow, bird, snake)", name)
			return
		}

		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("%s is not in (cow, bird, snake)", info)
			return
		}
	}
}
