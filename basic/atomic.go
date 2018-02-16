package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicint struct{
	value int
	lock sync.Mutex
}

func (a *atomicint)increment(){
	a.lock.Lock()
	defer a.lock.Unlock()
	(a.value)++
}

func (a *atomicint)get() int{
	a.lock.Lock()
	a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicint
	a.increment()
	go func(){
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
