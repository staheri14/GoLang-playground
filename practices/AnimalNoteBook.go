package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Write a program which allows the user to create a set of animals and to get information about those animals.
//Each animal has a name and can be either a cow, bird, or snake.
//With each command, the user can either create a new animal of one of the three types, or
//the user can request information about an animal that he/she has already created.
//Each animal has a unique name, defined by the user.
//Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
//The following table contains the three types of animals and their associated data.
//
//Animal	Food eaten	Locomtion method	Spoken sound
//cow	grass	walk	moo
//bird	worms	fly	peep
//snake	mice	slither	hsss
//Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
//Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
//Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.
//
//Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
//The second string is an arbitrary string which will be the name of the new animal.
//The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
//Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.
//
//Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal.
//The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
//Your program should process each query command by printing out the requested data.
//
//Define an interface type called Animal which describes the methods of an animal.
//Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
//The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
//Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command


type Animal interface {
	Eat()
	Move()
	Speak()
}
//===========================================
type Cow struct{
	name string
}
func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}
func  NewCow(name string) Cow{
	c:=  Cow{name}
	return c
}
//===========================================
type Bird struct{
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func  NewBird(name string) Bird{
	b:=  Bird{name}
	return b
}
//===========================================
type Snake struct{
	name string
}
func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func  NewSnake(name string) Snake{
	s:=  Snake{name}
	return s
}
//===========================================

func main() {
	AnimalNoteBook:=[]Animal{}

	//reads user input
	scanner := bufio.NewScanner(os.Stdin)



	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Split(input, " ")
		//the input should contain two parts
		if len(parts) != 3 {
			fmt.Println("Wrong number of arguments! try again")
			continue
		}

		query := strings.ToLower(parts[0])
		name := strings.ToLower(parts[1])
		switch query {
		case "newanimal": AnimalNoteBook=addAnimal(AnimalNoteBook,name,strings.ToLower(parts[2]))
		case "query": getAnimal(AnimalNoteBook,name,strings.ToLower(parts[2]))
		default:
			fmt.Println("Wrong command! try again")
		}

	}

}

func addAnimal(AnimalNoteBook []Animal,name, animalType string ) []Animal{
	//check the type of animal
	switch animalType {
	case "cow" :
		AnimalNoteBook = append(AnimalNoteBook, NewCow(name))
		fmt.Println("Created it!")
	case "bird" :
		AnimalNoteBook = append(AnimalNoteBook, NewBird(name))
		fmt.Println("Created it!")
	case "snake" :
		AnimalNoteBook = append(AnimalNoteBook, NewSnake(name))
		fmt.Println("Created it!")
	default:
		fmt.Println("Wrong name! try again")
	}
	return AnimalNoteBook

}

func getAnimal(AnimalNoteBook []Animal,name, action string ){

	for _,a :=range AnimalNoteBook{
		if c,ok:=a.(Cow); ok {
			if strings.Compare(c.name,name)==0{
				resolveAction(c,action)
				return
			}
		}
		if b,ok:=a.(Bird); ok {
			if strings.Compare(b.name,name)==0{
				resolveAction(b,action)
				return
			}
		}
		if s,ok:=a.(Snake); ok {
			if strings.Compare(s.name,name)==0{
				resolveAction(s,action)
				return
			}
		}

	}
	fmt.Println("Could not find the animal")
}

func resolveAction(animal Animal, action string){
	//resolve the request
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("The animal does not exist! try again")
	}

}
