package main

import (
	"fmt"
)

func change(s []int){
	s[0] = 100
}

func main(){
	arr := [...]int{0,1,2,3,4,5,6}
	s1 := arr[3:6]
	s2 := s1[2:4]
	fmt.Println("original slice")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)

	change(s1)
	change(s2)
	change(arr[:])
	fmt.Println("changed slice")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Printf("%v %d %d\n", s1, len(s1), cap(s1))
	fmt.Printf("%v %d %d\n", s2, len(s2), cap(s2))

	fmt.Println("array append")
	s3 := append(s1, 200)
	s4:=append(s2, 300)
	s4[0] = 99
	arr[5]=88
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println(s3)
	fmt.Println(s4)

	// delete s1 element
	fmt.Println("delete s1 element")
	fmt.Println(s1)
	copy(s1,s2)
	fmt.Println(s1,arr)
	copy(s1[1:],s1[2:])
	s1 = s1[:len(s1)-1]
	fmt.Println(s1)

	s3 = arr[:]
	fmt.Println(s3)
	s3 = append(s3, 0)
	copy(s3[4:], s3[3:])
	fmt.Println(s3)
}


