package main

import (
	"fmt"
)

type Worker struct {
	in chan int
	done chan bool
}

func worker(i int, in <-chan int, done chan<- bool){
	for n := range in{ // for loop way 1
			fmt.Printf("func worker %d, value = %c\n", i, n)
			go func(){done<-true}()
	//	}// end if
	}//end for loop

}

func createWorker(id int)Worker{
	c := Worker{
		in:make(chan int),
		done:make(chan bool),
	}
	go worker(id, c.in,c.done)
	return c
}
const  num=10
func chanDemo(){
	var workers [num]Worker
	for i:=0;i<num;i++{
		workers[i] = createWorker(i)
	}
	for i,worker:=range workers{
		worker.in <- 'a'+i
	}
	for i,worker:=range workers{
		worker.in<- 'A'+i
	}

	// wait for all of them
	for _, worker:=range workers{
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()

}
