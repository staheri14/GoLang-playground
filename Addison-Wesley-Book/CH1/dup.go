package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	inputs := make(map[string]int)
	lines := bufio.NewScanner(os.Stdin)
	for lines.Scan() {
		text := lines.Text()
		if strings.Compare(text, "exit") == 0 {
			break
		}
		inputs[text]++

	}

	for line, n := range inputs {
		fmt.Println("%d \t %s", n, line)
	}
}
