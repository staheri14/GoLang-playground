package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
//During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
//he program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
//The slice must grow in size to accommodate any number of integers which the user decides to enter.
//The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

func main() {

	var input int
	var err error
	var inputstring string
	intarr := make([]int, 0, 3)
	for {
		fmt.Println("Enter an integer (or x to exit): ")
		fmt.Scan(&inputstring)
		if strings.Compare(strings.ToLower(inputstring), "x") == 0 {
			return
		}

		//convert the input to integer
		if input, err = strconv.Atoi(inputstring); err != nil {
			fmt.Println("Enter a valid input")
			continue
		}

		//append it to the slice
		intarr = append(intarr, input)
		sort.Ints(intarr)
		fmt.Println(intarr)

		// if the max capacity is hit
		if len(intarr) == cap(intarr) {
			// double the capacity of the slice
			newintarr := make([]int, cap(intarr), (cap(intarr)+1)*2)
			copy(newintarr, intarr)
			intarr = newintarr
		}
	}

}
