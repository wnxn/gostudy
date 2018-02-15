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
	c1 := myGenerator()
	c2 := myGenerator()
	tm:=time.After(10*time.Second)
	w:=myCreateWorker(0)
	var values []int
	for{
		var activeWorker chan<-int
		var activeValue int
		if len(values)>0{
			activeWorker = w
			activeValue = values[0]
		}
		select{
		case n :=<-c1:
			values = append(values,n)
		case n :=<-c2:
			values = append(values,n)
		case activeWorker <-activeValue:
			fmt.Println(len(values))
			values = values[1:]
		case <-tm:
			fmt.Println("End of 10 seconds")
			return
		}
	}

}
