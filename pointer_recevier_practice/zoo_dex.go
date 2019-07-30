package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal structure to hold information about an animal
type Animal struct{ food, locomotion, noise string }

// Eat prints the food for an animal instance
func (a Animal) Eat() { fmt.Println(a.food) }

// Move prints the locomotion for an animal instance
func (a Animal) Move() { fmt.Println(a.locomotion) }

// Speak prints the noise for an animal instance
func (a Animal) Speak() { fmt.Println(a.noise) }

func main() {
	fmt.Println("Welcome. This program will get information about a predefined set of animals (cow, bird, snake).")
	fmt.Println("In order to make a request you should input 'animal property'. For example 'bird move'.")

	reader := bufio.NewReader(os.Stdin)

	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hss"},
	}

	for {
		fmt.Printf("> ")
		inputStr, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error: %v", err)
		}

		s := strings.Split(inputStr, " ")

		if animal, ok := animals[strings.ToLower(s[0])]; ok {
			property := strings.TrimSpace(strings.ToLower(s[1]))

			switch property {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Error: Property not found.")
			}
		} else {
			fmt.Println("Error: Animal not found.")
		}
	}
}
