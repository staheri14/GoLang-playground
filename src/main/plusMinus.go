package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	_ "io"
	"os"
	"strconv"
	"strings"
)

//Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with  places after the decimal.

// Complete the plusMinus function below.
func plusMinus(arr []int32) {

	var zcount, pcount, ncount float32

	for _, v := range arr {
		if v == 0 {
			zcount++
		} else if v < 0 {
			ncount++
		} else {
			pcount++
		}
	}
	var total float32 = float32(len(arr))
	fmt.Println(pcount / total)
	fmt.Println(ncount / total)
	fmt.Println(zcount / total)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}
