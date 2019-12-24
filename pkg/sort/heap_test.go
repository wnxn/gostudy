package sort

import (
	"fmt"
	"testing"
)

func getRow(rowIndex int) []int {
	if rowIndex == 1{
		return []int{1}
	}
	if rowIndex == 2{
		return []int{1,1}
	}
	a := getRow(rowIndex-1)
	res := make([]int, len(a)+1)
	res[0]=1
	for i:=1;i<len(a);i++{
		res[i]= a[i-1]+a[i]
	}
	res[len(a)]=1
	return res
}

func TestGet(t *testing.T){
	fmt.Println(getRow(3))
}
