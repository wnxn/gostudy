package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type people struct {
	name string
	age  int
}

func nilPeople() []people {
	return nil
}

func main() {
	fmt.Printf("%#v\n", []byte(s))
	//s := nilPeople()
	//fmt.Printf("len = %d, isnil %t", len(s), s==nil)
}
