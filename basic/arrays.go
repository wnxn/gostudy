package ch02

import "fmt"

func main(){
	var arr1 [5]int
	arr2 := [5]int{1,2,3,4,5}
	arr3 := [...]int{1,2,3}
	fmt.Println("printArr(arr1)")
	printArr(&arr1)
	fmt.Println("printArr(arr2)")
	printArr(&arr2)
	fmt.Println("print arr1 arr2 arr3")
	fmt.Println(arr1, arr2,arr3)
}

func printArr(arr *[5]int){
	arr[0]=100
	for i,v := range arr{
		fmt.Println(i,v)
	}
}
