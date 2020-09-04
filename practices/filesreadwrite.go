package main

import (
	"fmt"
	"io/ioutil"
	"main/helper"
	"os"
)

func main() {

	ioutilread("hello.txt")
	ioutilwrite("bye.txt", "Bye!")

	fmt.Println()
	osread("hello.txt")
	oswrite("output.txt", "Bye!")
}

func osread(filename string) {
	//first open a file
	f, err := os.Open(filename)
	helper.CheckError(err)

	//read the file
	barr := make([]byte, 13)
	f.Read(barr)
	fmt.Println("The content of the input file (", filename, ") is: ", string(barr))

	//close the file
	f.Close()
}

func oswrite(filename string, out string) {
	//first open a file
	f, err := os.Create(filename)
	helper.CheckError(err)

	//Write into the file
	f.WriteString(out)
	fmt.Println("The content (", out, ") is copied into the ", filename)

	//close the file
	f.Close()
}

func ioutilwrite(filename string, out string) {
	//
	outb := []byte(out)
	ioutil.WriteFile(filename, outb, 0777)
	fmt.Println("The content (", out, ") is copied into the ", filename)

}

func ioutilread(filename string) {
	in, err := ioutil.ReadFile(filename)
	helper.CheckError(err)
	//in is a byte array
	fmt.Println("The byte representation of the input file (", filename, ") is: ", in)
	fmt.Println("The string representation of the input file (", filename, ") is: ", string(in))

}
