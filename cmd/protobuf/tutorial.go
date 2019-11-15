package main

import "fmt"

func main(){
	var a []int
	a = append(a,2)
	if a == nil{
		fmt.Println("==nil")
	}else{
		fmt.Println("!=nil")
	}
}