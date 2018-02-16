package main

import (
	"fmt"
	"time"
	//"reflect"
)

func worker(i int, c <-chan int){
	//for n := range c{ // for loop way 1
	for {				// for loop way 2
	//	if v,ok:=<-c;ok {	// get chan value way 1
			v:= <-c				// get chan value way 2
			fmt.Printf("func worker %d, value = %c\n", i, v)
	//	}// end if
	}//end for loop

}

func createWorker(id int)chan<-int{
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo(){
	var channels [10]chan<-int
	for i:=0;i<10;i++{
		channels[i] = createWorker(i)
	}
	for i:=0;i<10;i++{
		channels[i] <- 'a'+i
	}
	for i:=0;i<10;i++{
		channels[i]<- 'A'+i
	}
	time.Sleep(time.Millisecond*10)
}

func bufferedChannel(){
	c:=make(chan int, 3)
	go worker(0,c)
	c <-'A'+1
	c<-'A'+2
	c<-'A'+3
	c<-'A'+4
	time.Sleep(time.Millisecond)
}

func channelClose(){
	c:=make(chan int)
	go worker(0,c)
	c <-'A'+1
	c<-'A'+2
	c<-'A'+3
	c<-'A'+4
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close")
	channelClose()

}
