package main

import (
	"fmt"
	"sync"
)

//the global  variable shared between routines
var i int = 0
var wg sync.WaitGroup

func inc() {
	i = i + 1
	wg.Done()
}

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}
