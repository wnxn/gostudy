package main

import (
	"fmt"
)

func main(){
	var a []int
	 a = append(a,[]int{1,2,3}...)
	fmt.Println(a)

	 for _, v:=range a{
	 	a = append(a,v)
	 }
	 fmt.Println(a)
}
