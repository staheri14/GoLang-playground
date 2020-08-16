package main

import (
	"fmt"
	"io/ioutil"
	"main/helper"
)

func main(){
	in,err:=ioutil.ReadFile("hello.txt")
	helper.CheckError(err)
	//in is a byte array
	fmt.Println("The byte representation of the input file is: " ,in)
	fmt.Println("The string representation of the input file is: " ,string(in))


	//
	out:=[]byte("Bye!")
	ioutil.WriteFile("bye.txt", out,0777)

}
