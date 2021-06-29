// Echo its command line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// normal for loop
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// with range
	s = ""
	sep = " "
	for _, val := range os.Args[1:] {
		s += sep + val
	}
	fmt.Println(s)

	// print the first agr as well
	fmt.Println(strings.Join(os.Args[:], " "))

	// with range and args index
	s = ""
	sep = " "
	for i, val := range os.Args[1:] {
		s += sep + "[" + strconv.Itoa(i) + "]:" + sep + val
	}
	fmt.Println(s)
}
