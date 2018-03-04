package main

import (
	"fmt"
	"reflect"
)

func main() {
	var circle float64 = 3.14

	value := reflect.ValueOf(circle)
	fmt.Println("Reflect: value=",value)
	fmt.Println("Reflect can set ", value.CanSet())

	valueptr:=reflect.ValueOf(&circle)
	value3 :=valueptr.Elem()
	value3.SetFloat(2.2)
	fmt.Println("Reflect: value=",valueptr)
	fmt.Println("Reflect can set ", valueptr.CanSet())

}
