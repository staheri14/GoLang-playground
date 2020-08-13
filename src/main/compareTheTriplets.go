package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//Alice and Bob each created one problem for HackerRank. A reviewer rates the two challenges, awarding points on a scale from 1 to 100 for three categories: problem clarity, originality, and difficulty.
//
//The rating for Alice's challenge is the triplet a = (a[0], a[1], a[2]), and the rating for Bob's challenge is the triplet b = (b[0], b[1], b[2]).
//
//The task is to find their comparison points by comparing a[0] with b[0], a[1] with b[1], and a[2] with b[2].
//
//If a[i] > b[i], then Alice is awarded 1 point.
//If a[i] < b[i], then Bob is awarded 1 point.
//If a[i] = b[i], then neither person receives a point.
//Comparison points is the total points a person earned.
//
//Given a and b, determine their respective comparison points
// Input Format
//
//The first line contains 3 space-separated integers, a[0], a[1], and a[2], the respective values in triplet a.
//The second line contains 3 space-separated integers, b[0], b[1], and b[2], the respective values in triplet b.
//
//Constraints
//
//1 ≤ a[i] ≤ 100
//1 ≤ b[i] ≤ 100

//computes the scores of comparison of arrays a and b
func compareTriplets(a []int32, b []int32) []int32 {

	scores := []int32{0, 0}
	for i, v := range a {
		if v > b[i] {
			scores[0] = scores[0] + 1
		}
		if v < b[i] {
			scores[1] = scores[1] + 1
		}
	}
	return scores

}

func main() {
	//creates a reader with the given buffer size in bytes
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//creates the output file
	stdout, err := os.Create("compareTheTriplets.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	//reads Alice's input ============================================
	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	//keeps Alice's input
	var a []int32

	for i := 0; i < 3; i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}
	//reads Bob's input ============================================

	bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	//keeps Bob's input
	var b []int32

	for i := 0; i < 3; i++ {
		bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	//compares the inputs ========================================
	result := compareTriplets(a, b)

	//prints the result
	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}



