package main

import (
	"fmt"
)

func test(a interface{}){
	if v,ok:=a.(string); ok{
		fmt.Printf("string %s\n",v)
	}
	if v,ok:=a.(int);ok{
		fmt.Printf("value %d\n",v)
	}
	switch v:=a.(type) {
	case string:
		fmt.Println(v," is string")
	case int:
		fmt.Println(v," is int")
	}
}

func main() {
	test("H")
	test(1)
}
