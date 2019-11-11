package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	arr := [10]int{}
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				arr[ii]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(arr)
}
