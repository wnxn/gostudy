package main

import (
	"github.com/wnxn/gostudy/algorithm/linklist"
	"fmt"
)

func main() {
	arr :=[]int{1,2,3,4,5}
	link := linklist.CreateRecursive(arr)
	linklist.Print(link)
	fmt.Println()
	result:=linklist.ReverseLinklist2(link)
	linklist.Print(result)
}
