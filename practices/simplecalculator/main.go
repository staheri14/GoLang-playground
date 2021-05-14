package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fr := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number:")
	input1, _:= fr.ReadString('\n')

	fmt.Println("Enter another number:")
	input2, _ := fr.ReadString('\n')

	fmt.Println("Enter an operator:")
	op, _ := fr.ReadString('\n')
	op = strings.TrimSpace(op)
	v1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err1 != nil {
		panic(err1)
	}
	v2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err2 != nil {
		panic(err1)
	}

	var res float64 = 0
	switch strings.TrimSpace(op) {
	case "+":
		res = v1 + v2
		break
	case "-":
		res = v1 - v2
		break
	case "/":
		if v2!=0 {res = v1/v2} else { panic("denominator cannot be zero")}
		break
	case "*":
		res = v1 * v2
		break
	default:
		panic("Invalid operator")
	}

	fmt.Printf("%v %v %v = %v \n", v1, op, v2, res)


}
