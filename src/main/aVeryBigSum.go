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
	"io"
	_ "io"
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

	// get the value of the environment variable OUTPUT_PATH
	stdout, err := os.Create("aVeryBigSum.txt")
	checkErr(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLn(reader), 10, 64)
	checkErr(err)

	arTemp := strings.Split(readLn(reader), " ")

	var ar []int64

	for i := 0; i < int(arCount); i++ {
		arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkErr(err)
		ar = append(ar, arItem)
	}

	result := aVeryBigSum(ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readLn(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}