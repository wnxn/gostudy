package main

import "fmt"

type people struct{
	name string
	age int
}

func nilPeople()[]people{
	return nil
}

func main() {
	s := nilPeople()
	fmt.Printf("len = %d, isnil %t", len(s), s==nil)
}
