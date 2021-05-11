package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fileReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your text: ")
	input, _ := fileReader.ReadString('\n')
	fmt.Println("Entered: ", input)

	fmt.Println("Enter as number: ")
	aNumber, _ := fileReader.ReadString('\n')
	aFloat, error := strconv.ParseFloat(strings.TrimSpace(aNumber), 64)
	if error != nil{
		panic(error)
	}
	fmt.Println("Entered number: ", aFloat)
}
