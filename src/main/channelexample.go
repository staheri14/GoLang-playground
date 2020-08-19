package main

import "fmt"

func main(){
	c:= make(chan int)
	go prod(1,2,c)
	go prod(3,4,c)
	a:=<-c
	b:=<-c
	fmt.Println("The product of 1,2,3,4 is ",a*b)

}

func prod(a,b int,c chan int ){
	c<-a*b
}

