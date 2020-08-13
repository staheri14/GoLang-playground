package main

//To calculate and print the sum of the elements in an array, keeping in mind that some of those integers may be quite large.
// Complete the aVeryBigSum function below.
//Input Format
//The first line of the input consists of an integer .
//The next line contains  space-separated integers contained in the array.
//
//Output Format
//Return the integer sum of the elements in the array.

import (
	"bufio"
	"fmt"
	_ "io"
	"main/helper"
	"os"
	"strconv"
	"strings"
)

func aVeryBigSum(ar []int64) int64 {

	var sum int64
	for _, v := range ar {
		sum = sum + v
	}
	return sum
}

func main() {
	//a reader with at least 2^20 bytes buffer size
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// prepare the output file
	stdout, err := os.Create("aVeryBigSum.txt")
	helper.CheckError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	fmt.Println("Enter the size of the array")
	//gets the array size
	arCount, err := strconv.ParseInt(helper.ReadLine(reader), 10, 64)
	helper.CheckError(err)

	fmt.Println("Enter the space-separated integers")
	//gets the space-separated integers
	arTemp := strings.Split(helper.ReadLine(reader), " ")

	var ar []int64

	//retrieves the integers from the arTemp
	for i := 0; i < int(arCount); i++ {
		arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
		helper.CheckError(err)
		ar = append(ar, arItem)
	}

	result := aVeryBigSum(ar)

	fmt.Fprintf(writer, "%d\n", result)

	// writes the output
	writer.Flush()
}

