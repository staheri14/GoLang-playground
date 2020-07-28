package main

import (
"fmt"
"strconv"
)

func main(){
	var x float64
	fmt.Print("Enter a floating point number: \n")
	fmt.Scan(&x)
	converted_value:=strconv.FormatFloat(x,'f',0,64)
	fmt.Print(converted_value)

}
