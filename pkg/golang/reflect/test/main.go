package main

import (
	"fmt"
	"reflect"
)

func main() {
	x:=2
	a:=reflect.ValueOf(&x)
	d := a.Elem()
	
}
