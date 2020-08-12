package main
//which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake.
//Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal:
//1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
//The following table contains the three animals and their associated data which should be hard-coded into your program.
//
//Animal	Food eaten	Locomotion method	Spoken sound
//cow	grass	walk	moo
//bird	worms	fly	peep
//snake	mice	slither	hsss

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, sound string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	//defines the animals
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	//reads user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your request:")

	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Split(input, " ")
		//the input should contain two parts
		if len(parts) != 2 {
			fmt.Println("Wrong input! try again")
			continue
		}
		name := strings.ToLower(parts[0])
		request := strings.ToLower(parts[1])


		//check the type of animal
		var animal Animal
		switch name {
		case "cow" : animal=cow
		case "bird" : animal=bird
		case "snake" : animal=snake
		default:
			fmt.Println("Wrong name! try again")
			continue

		}

		//resolve the request
		switch request {
		case "eat": animal.Eat()
		case "move": animal.Move()
		case "speak": animal.Speak()
		default:
			fmt.Println("Wrong request! try again")
			continue
		}

		fmt.Println("Enter your request:")

	}

}

