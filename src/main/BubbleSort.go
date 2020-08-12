package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
//The program should print the integers out on one line, in sorted order, from least to greatest.
//Use your favorite search tool to find a description of how the bubble sort algorithm works.
//
//As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
//The BubbleSort() function should modify the slice so that the elements are in sorted order.
//
//A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
//You should write a Swap() function which performs this operation.
//Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
//The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

func main() {

	max_length := 10
	array := make([]int, 0, max_length)
	//populate the array with the user inputs
	for i := 0; i < max_length; i++ {
		fmt.Println("Type Q to end or type an Integer:")
		var v string
		fmt.Scan(&v)
		//break if user entes Q
		if strings.Compare(strings.ToLower(v), "q") == 0 {
			break
		}
		vi, _ := strconv.Atoi(v)
		array = append(array, vi)
	}
	fmt.Println(array)
	BubbleSort(array)
	fmt.Println("Sorted array: ", array)

}

//Sorts the given slice
func BubbleSort(array []int) {

	for i := range array {
		sci := array[0 : len(array)-i-1]
		for j := range sci {
			//move the max toward the end of the array
			if array[j] > array[j+1] {
				swap(array, j)
			}
		}
	}
}

//Swaps the two elements corresponding to index and index+1
func swap(array []int, index int) {
	temp := array[index]
	array[index] = array[index+1]
	array[index+1] = temp
}
