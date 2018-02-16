package main

import "fmt"

func domain(){
	defer fmt.Printf("outter defer\n")
	value:=2
	for i:=0;i<2;i++{
		defer fmt.Printf("inner defer\n")
		fmt.Printf("value is %d\n", i)
	}
	fmt.Printf("value is %d\n", value)
}

func main() {
	domain()
}
