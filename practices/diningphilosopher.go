package main

import (
	"fmt"
	"sync"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

type Chopstick struct{sync.Mutex}

type Philosopher struct {
	Index int
	RightChopstcik, leftChopstick Chopstick
}

func (philosopher Philosopher) eat( host chan int, w *sync.WaitGroup){
	defer w.Done()
	for i:=0;i<3;i++{

		<-host //get permission from the host
		philosopher.leftChopstick.Lock()
		philosopher.RightChopstcik.Lock()
		fmt.Println("Starting to eat ", philosopher.Index)
		philosopher.leftChopstick.Unlock()
		philosopher.RightChopstcik.Unlock()
		fmt.Println("Finishing eat ", philosopher.Index)
		host<-1 //release the permission
	}
}


func main(){
	Host:=make(chan int,2)
	Host<-1
	Host<-1

	w:=&sync.WaitGroup{}
	w.Add(5)

	chopsticks:=make([]Chopstick,5)
	philosophers:=make([]Philosopher,5)
	for i:=0;i<5;i++{
		philosophers[i].Index=i
		philosophers[i].leftChopstick=chopsticks[i%5]
		philosophers[i].leftChopstick=chopsticks[(i+1)%5]
	}
	for i:=0;i<5;i++{
		go philosophers[i].eat(Host,w)
	}
	w.Wait()
}
