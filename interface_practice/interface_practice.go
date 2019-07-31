package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal structure
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow type
type Cow struct {
	food, locomotion, noise string
}

//Bird type
type Bird struct {
	food, locomotion, noise string
}

//Snake type
type Snake struct {
	food, locomotion, noise string
}

//Eat - Cow
func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

//Eat - Bird
func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

//Eat - Snake
func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

//Move - Cow
func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

//Move - Bird
func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

//Move - Snake
func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

//Speak - Cow
func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

//Speak - Bird
func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

//Speak - Snake
func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}

func main() {

	var nameType map[string]string

	nameType = make(map[string]string)

	for {
		fmt.Printf(">please enter the request -> name & infomation: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()

		//fmt.Printf("request is: %s\n", request)

		if len(strings.Split(request, " ")) != 3 {
			fmt.Println("Command input number is not equal to 3")
			continue
		}

		command := strings.Split(request, " ")[0]
		fmt.Printf("command: %s\n", command)

		switch command {
		case "newanimal":

			name := strings.Split(request, " ")[1]
			animalType := strings.Split(request, " ")[2]

			fmt.Printf("name: %s\n", name)
			fmt.Printf("animalType: %s\n", animalType)

			//adding name checking
			nameType[name] = animalType
		case "query":

			name := strings.Split(request, " ")[1]
			info := strings.Split(request, " ")[2]

			fmt.Printf("name: %s\n", name)
			fmt.Printf("info: %s\n", info)

			switch nameType[name] {
			case "cow":
				cow := Cow{"grass", "walk", "moo"}
				switch info {
				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			case "bird":
				bird := Bird{"worms", "fly", "peep"}
				switch info {
				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			case "snake":
				snake := Snake{"mice", "slither", "hsss"}
				switch info {
				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			default:
				fmt.Printf("%s is not in (cow, bird, snake), please try again\n", nameType[name])
				continue
			}
		default:
			fmt.Printf("%s is not in (newanimal, query), please try again\n", command)
			continue
		}
	}
}
