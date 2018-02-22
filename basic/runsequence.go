package main

import (
	"fmt"
)

func messageHello() string{
	fmt.Println("Run messageHello() string domain")
	return "Hello"
}



func init(){
	fmt.Printf("init func\n")
}

func main() {
	ch := make(chan bool,1)
	fmt.Printf("cap=%d\n",cap(ch))
	nonamefunc := func(){
		fmt.Printf("This is a annoymous func\n")
	}
	nonamefunc()

	func(){
		fmt.Printf("This is another annoymous func\n")
	}()
	fmt.Println(str)
}

var str = messageHello()