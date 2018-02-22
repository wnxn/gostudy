package main

import (
	"fmt"
)

func makesilce(value int) []string{
	ret := make([]string, value)
	for i:=0;i<value;i++{
		ret[i]= fmt.Sprintf("No.%d, is %d", i, i)
	}
	return ret
}

func closure(){
	strslice := makesilce(20)
	for _,v:=range strslice{
		go func(){
			fmt.Println(v)
		}()
	}
}

func main() {
	closure()
}
