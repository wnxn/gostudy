package main

import (
	"fmt"
	"time"
	"runtime"
)
func send(ch chan int){
	i:=1
	ch<-i
	for{
		fmt.Printf("send %d\n",i)
		i = <-ch
		i++
		ch<-i
	}

}

func recv(ch chan int){
	for{
		i:=<-ch
		fmt.Printf("recv %d\n",i)
		ch<-i
	}

}
func main(){
	runtime.GOMAXPROCS(4)
	tm:=time.After(time.Second)
	ch1:=make(chan int)
	go send(ch1)
	go recv(ch1)
	<-tm
}
