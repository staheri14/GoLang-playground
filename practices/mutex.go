package main

import (
	"fmt"
	"sync"
)

//the global  variable shared between routines
var j int = 0
var wgr sync.WaitGroup

var mu sync.Mutex

func incr() {
	mu.Lock()
	j = j + 1
	mu.Unlock()
	wgr.Done()

}

func main() {
	wgr.Add(2)
	go incr()
	go incr()
	wgr.Wait()
	fmt.Println(j)
}
