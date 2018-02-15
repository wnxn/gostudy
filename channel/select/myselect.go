package main

import (
	"fmt"
	"math/rand"
	"time"
)

func myWorker(i int, c <-chan int) {
	for n := range c { // for loop way 1
		time.Sleep(time.Second)
		fmt.Printf("func worker %d, value = %d\n", i, n)
	} //end for loop
}

func myCreateWorker(id int) chan<- int {
	c := make(chan int)
	go myWorker(id, c)
	return c
}

func myGenerator() <-chan int {
	c := make(chan int)
	value := 0
	go func() {
		for {
			c <- value
			value++
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))
		}
	}()
	return c
}

func main() {
	var c1, c2 = myGenerator(), myGenerator()
	var worker = myCreateWorker(0)
	var values []int
	n := 0
	tm:=time.After(10*time.Second)
	tick:=time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values)>0 {
			activeWorker = worker
			activeValue=values[0]
		}
		select {
		case n = <-c1:
			values = append(values,n)
		case n = <-c2:
			values = append(values,n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800*time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
