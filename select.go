package main

import "fmt"

func main(){
	c1:=make(chan int)
	c2:=make(chan int)

	go product(10,12,c1)
	go product(10,12,c2)

	select {
	case res:=<-c1:
		fmt.Println(res, "Comes from c1")
	case res:=<-c2:
		fmt.Println(res, "Comes from c2")
	}

}

func product(a,b int, c chan int){
	c<-a*b
}


