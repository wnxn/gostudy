package main

import (
	"fmt"
)

func tryRecover(){
	defer func(){
		r :=recover()
		if err,ok:=r.(error);ok{
			fmt.Println("Error occure: ", err.Error())
		}else{
			panic(fmt.Sprintf("I don't know what to do %T %v", r,r))
		}
	}()
	// error 1
	//panic(errors.New("this is an error"))

	// error 2
	//b:=0
	//a:=5/b
	//fmt.Println(a)

	//error 3
	panic(123)
}

func main() {
	tryRecover()
}
