package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i:=0; i < 10; i++{
		fmt.Printf("%X,%X\n",&i, &a[0])
		go func(i int){
			for{
				a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
