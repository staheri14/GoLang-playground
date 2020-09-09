package main

import (
	"fmt"
	"sync"
)

//Write two goroutines which have a race condition when executed concurrently.
//Explain what the race condition is and how it can occur.


// Race conditions occur due to communication between two concurrent go routines. More specifically, when two routines share a value.
// and the routines attempt write and read on the shared data concurrently, due to which the order by which
// they access the data will affect the output and will differ for different inter-leavings
func main(){
	fmt.Println("Race conditions occur due to communication between two concurrent go routines. More specifically, when two routines share a value")
	fmt.Println("and the routines attempt write and read on the shared data concurrently, thus the order by which")
	fmt.Println("the routines access the data will affect the output and will differ for different inter-leavings")
	fmt.Println("In the example below, the expected output is: 1,...,20 on one line and 21,...,40 on another line but due to the race condition the output below may differ (please run it multiple times to see the difference)")
	 x:=0
	 w:=sync.WaitGroup{}
	 w.Add(2)
	 go routineA(&x,&w)
	 go routineB(&x,&w)
	 w.Wait()
}

func routineA(x *int, w *sync.WaitGroup ){
	//fmt.Println("Routine A: ")
	for i:=0;i<20;i++{
		*x=(*x)+1
		fmt.Print(" ",*x)
	}
	fmt.Println()
	 w.Done()
}

func routineB(x *int, w *sync.WaitGroup ){
	//fmt.Println("Routine B: ")
	for i:=0;i<20;i++{
		*x=(*x)+1
		fmt.Print(" ",*x)
	}
	fmt.Println()
	w.Done()
}

