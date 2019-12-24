package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var x []int
	a := sync.WaitGroup{}
	a.Add(2)
	rand.Seed(time.Now().Unix())
	go func() {
		defer a.Done()
		d := time.Millisecond*time.Duration(rand.Int()%100)
		fmt.Printf("1: sleep %s\n", d.String())
		time.Sleep(d)
		fmt.Printf("1: assigned\n")
		x= make([]int, 100)
	}()
	go func() {
		defer a.Done()
		d := time.Millisecond*time.Duration(rand.Int()%100)
		fmt.Printf("2: sleep %s\n", d.String())
		time.Sleep(d)
		fmt.Printf("2: assigned\n")
		x = make([]int, 10000)
	}()
	a.Wait()
	fmt.Println(len(x))
}
