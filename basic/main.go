package main

import "fmt"

func messageHello() string{
	fmt.Println("Run messageHello() string")
	return "Hello"
}

var str = messageHello()

func main() {
	nonamefunc := func(){
		fmt.Printf("This is a annoymous func\n")
	}
	nonamefunc()

	func(){
		fmt.Printf("This is another annoymous func\n")
	}()
	fmt.Println(str)
}
