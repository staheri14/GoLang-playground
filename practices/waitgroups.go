package main

// this  program uses WaitGroup object to handle dependency between two Go routines i.e., the foo and the main routines
import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	//acknowledge the execution at the end
	defer wg.Done()

	fmt.Println("foo routine")
}
func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	//waits for foo to finish
	wg.Wait()
	fmt.Println("The main routine")
}
