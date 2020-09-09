package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

//Write a program to sort an array of integers.
//The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
//Each partition should be of approximately equal size.
//Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
//
//The program should prompt the user to input a series of integers.
//Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
//When sorting is complete, the main goroutine should print the entire sorted list.


func main() {


	fmt.Println("Insert all your input integers space separated on the same line:")
	br:=bufio.NewReader(os.Stdin)
	v,_, _ :=br.ReadLine()
	line:=string(v)
	numbersstr:=strings.Fields(line)

	//converts the input to int values
	array := make([]int, 0,len(numbersstr))
	for _,v := range numbersstr {
		vi,_:=strconv.Atoi(v)
		array= append(array, vi)
	}

	var q int
	q=len(numbersstr)/4

	w:=&sync.WaitGroup{}
	w.Add(4)

	go Sort(array[:q],w)
	go Sort(array[q:2*q],w)
	go Sort(array[2*q:3*q],w)
	go Sort(array[3*q:],w)

	w.Wait()
	c1:=merge(array[:q],array[q:2*q])
	c2:=merge(array[2*q:3*q],array[3*q:])
	c3:=merge(c1,c2)

	fmt.Println("Sorted array: ", c3)

}

// Sort sorts the given slice
func Sort(array []int, w *sync.WaitGroup) {
	defer w.Done()
	sort.Ints(array)
	defer fmt.Println(array)
}


func merge (p1,p2 []int) []int{
	sortedarr:=make([]int,0, len(p1)+len(p2))

	for len(p1)!=0 || len(p2)!=0{
		if len(p1)!=0 && len(p2)==0{
			sortedarr = append(sortedarr, p1[0])
			p1=p1[1:]
			continue
		}
		if len(p1)==0 && len(p2)!=0 {
			sortedarr = append(sortedarr, p2[0])
			p2=p2[1:]
			continue
		}

		if p1[0]<p2[0]{
			sortedarr=append(sortedarr,p1[0])
			p1=p1[1:]
		}else {
			sortedarr=append(sortedarr,p2[0])
			p2=p2[1:]
		}
	}

	return sortedarr
}