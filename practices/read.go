package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Write a program which reads information from a file and represents it in a slice of structs.
//Assume that there is a text file which contains a series of names.
//Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
//
//Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
//Each field will be a string of size 20 (characters).
//
//Your program should prompt the user for the name of the text file.
//Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
//Each struct created will be added to a slice, and after all lines have been read from the file,
//your program will have a slice containing one struct for each line in the file.
//After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.


type Person struct {
	Fname, Lname  string
}

type People []Person

func main(){
	var filename string
	var f *os.File
	var err error

	for {
		fmt.Println("Enter your file name")
		fmt.Scan(&filename)

		//open the file
		f, err = os.Open(filename)
		if err != nil {
			fmt.Println("The file does not exist")
			continue
		}
		break
	}

	// populate the list of people
	var list People
	scanner:=bufio.NewScanner(f)
	for scanner.Scan() {
		// read the next line
		line:=scanner.Text()
		// separate the first and last name in each line
		tokens:=strings.Split(line, " ")

		//check whether there are first and last name
		if len(tokens)<2 {
			continue
		}
		// check whether they are less than 20 chars
		if len(tokens[0])>20 || len(tokens[1])>20{
			continue
		}

		p:=Person{tokens[0],tokens[1]}
		list = append(list, p)
	}

	// printing the slice into the console
	for _,x := range list{
		fmt.Println( x.Fname, " ", x.Lname)
	}

}
