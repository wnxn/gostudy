package main

import (
	"fmt"
)

func ch8(){
	ch:=make(chan int)
	for i:=0;i<10;i++{
		select {
		case x:=<-ch:
			fmt.Println(x)
		case ch<-i:
		default:
			fmt.Println("default")
		}
	}
}

func oneroutine()string{
	ch :=make(chan string)
	go func(){ch <-"go func 1"}()
	go func(){ch <-"go func 2"}()
	go func(){ch <-"go func 3"}()
	return <-ch
}

func main() {
	ch8()
}
