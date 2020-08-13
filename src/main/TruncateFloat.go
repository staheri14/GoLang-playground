package main

/**
Write a program which prompts the user to enter a floating point number and prints the integer which is
a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/
import (
	"fmt"
	"strconv"
)

func main() {
	var x float64
	fmt.Print("Enter a floating point number: \n")
	fmt.Scan(&x)
	converted_value := strconv.FormatFloat(x, 'f', 0, 64)
	fmt.Print(converted_value)

}
