package main

import (
	"fmt"
	"sync"
)


func foo(wg *sync.WaitGroup){
	//acknowledge the execution at the end
	defer wg.Done()

	fmt.Println("foo routine")
}
func main(){

	var wg sync.WaitGroup
	wg.Add(1)
	foo(&wg)
	//waits for foo to finish
	wg.Wait()
	fmt.Println("The main routine")
}
