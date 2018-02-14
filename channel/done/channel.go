package main

import (
	"fmt"
	"sync"
)


type Worker struct {
	in chan int
	done func()
}

func worker(i int, c Worker){
	for n := range c.in{ // for loop way 1
			fmt.Printf("func worker %d, value = %c\n", i, n)
			c.done()
			//go func(){done<-true}()
	//	}// end if
	}//end for loop

}

func createWorker(id int, w *sync.WaitGroup)Worker{
	c := Worker{
		in:make(chan int),
		done: w.Done,
	}
	go worker(id, c)
	return c
}
const  num=20
func chanDemo(){
	var workers [num]Worker
	var wg sync.WaitGroup
	for i:=0;i<num;i++{
		workers[i] = createWorker(i,&wg)
	}
	for i,worker:=range workers{
		wg.Add(1)
		worker.in <- 'a'+i

	}
	for i,worker:=range workers{
		wg.Add(1)
		worker.in<- 'A'+i

	}
	wg.Wait()
}

func main() {
	chanDemo()

}
