package main

import (
	"fmt"

)

func sendrecvoneroutine(){
	c := make(chan int,1)
	c<-1
	n:=<-c
	fmt.Printf("sendrecvoneroutine %d\n",n)
	c<-n+1
}

func main() {
	sendrecvoneroutine()
}
