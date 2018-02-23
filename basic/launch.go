package main

import (
	"time"
	"fmt"
	"os"
)
func abort(ret chan struct{}){
	os.Stdin.Read(make([]byte,1))
	ret <- struct{}{}
}

func main() {
	sig := make(chan struct{})
	go abort(sig)
	tm :=time.Tick(time.Second)
	for i:=1;i<10;i++{
		select{
		case<-tm:
			fmt.Printf("%d second left\n", i)
		case <-sig:
			fmt.Println("abort")
			return
		}
	}
}
