package main

import (
	"fmt"
	"sync"
	"time"
)

func sendint(ch chan<-int, wg *sync.WaitGroup){
	defer close(ch)
	defer wg.Done()
	tm :=time.After(time.Second*10)
	value :=0
	for{
		select {
		case ch<-value:
			fmt.Printf("send int %d\n", value)
		case <-tm:
			return
		}
		value++
	}
}

func recvint(ch <-chan int, wg *sync.WaitGroup){
	defer wg.Done()
	ok:=false
	for{
		var value int
		if value,ok=<-ch; ok{
			time.Sleep(time.Second)
			fmt.Printf("recv int %d\n", value)
		}else{
			return
		}
	}
}

func main() {
	ch :=make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go sendint(ch,&wg)
	go recvint(ch,&wg)
	wg.Wait()
	fmt.Printf("main return normally\n")
}
