package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countLines(f *os.File, counts map[string]int) {
	lines := bufio.NewScanner(f)
	for lines.Scan() {
		text := lines.Text()
		if strings.Compare(text, "exit") == 0 {
			break
		}
		counts[text]++

	}
}

func main() {

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fname := range files {
			file, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stdin, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
		}
	}
	for line, n := range counts {
		fmt.Println("%d \t %s", n, line)
	}
}
