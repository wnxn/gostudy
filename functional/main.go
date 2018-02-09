package main

import (
	"fmt"
)

func adder() func (int)int{
	fmt.Println("func(int)int")
	sum := 0
	return func(v int)int{
		fmt.Println("func(v int)int")
		sum +=v
		return sum
	}
}

func main() {
	a:=adder()
	fmt.Printf("a.(type)=%T\n", a)
	for i:=0;i<10;i++ {
		fmt.Println(a(i))
	}
}
