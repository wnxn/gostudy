package main

import (
	"runtime"
	"sync"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Number of CPU: ",runtime.NumCPU())
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func(){
		defer wg.Done()

		for i:=0; i<3;i++{
			for j:='a'-1; j<='z';j++{
				time.Sleep(time.Millisecond*10)
				fmt.Printf("%c ", j)
			}
		}
	}()
	go func(){
		defer wg.Done()

		for i:=0; i<3;i++{
			for j:='A'-1; j<='Z';j++{
				time.Sleep(time.Millisecond*10)
				fmt.Printf("%c ", j)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")

}
