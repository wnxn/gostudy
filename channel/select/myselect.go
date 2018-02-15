package main

import (
	"fmt"
	"time"
//	"math/rand"
)

func myGenerator() <-chan int{
	c := make(chan int)
	value := 0
	go func(){
		for{
			c <- value
			value++
			fmt.Printf("myGenerator %d\n", value)
	//		time.Sleep(time.Millisecond *time.Duration(rand.Int()%1500))
		}

	}()
	return c
}

func main() {
	c1 := myGenerator()
	c2 := myGenerator()
	tm:=time.After(10*time.Second)
	num:=0
	for{
		select{
		case n :=<-c1:
	//		time.Sleep(time.Millisecond*2)
			fmt.Printf("No. %d: Received from c1 %d\n", num, n)
			num++
		case n :=<-c2:

			fmt.Printf("No. %d: Received from c2 %d\n", num, n)
			time.Sleep(time.Millisecond)
			num++
		case <-tm:
			fmt.Println("End of 10 seconds")
			return
		}
	}

}
