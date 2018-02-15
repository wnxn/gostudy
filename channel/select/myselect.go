package main

import (
	"fmt"
	"time"
	"math/rand"
)

func myWorker(i int, c <-chan int){
	for n := range c{ // for loop way 1
		fmt.Printf("func worker %d, value = %d\n", i, n)
	}//end for loop

}

func myCreateWorker(id int)chan<-int{
	c := make(chan int)
	go myWorker(id, c)
	return c
}


func myGenerator() <-chan int{
	c := make(chan int)
	value := 0
	go func(){
		for{
			c <- value
			value++
			time.Sleep(time.Millisecond *time.Duration(rand.Intn(1500)))
		}

	}()
	return c
}

func main() {
	var c1, c2 = myGenerator(),myGenerator()
	var worker = myCreateWorker(0)

	n:=0
	hasValue:=false
	for {
		var activeWorker chan<- int
		if hasValue{
			activeWorker = worker
		}

		select {
		case n = <-c1:
			hasValue=true
		case n = <-c2:
			hasValue=true
		case activeWorker <- n:
			hasValue = false
		}
	}


}
