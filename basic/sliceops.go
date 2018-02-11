package main
import (
	"fmt"
)

func create(len int)( []int){
	var ret []int
	for i:=0;i<len;i++{
		printSlice(ret)
		ret = append(ret, i*2 +1)
	}
	return ret
}

func printSlice(s []int){
	fmt.Printf("%v len = %d cap = %d\n", s, len(s), cap(s))
}
func main(){
	s1 :=create(20)
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))

	s3 := s1[:32]
	fmt.Println(copy(s3, s1))
	printSlice(s3)

	s3 = append(s3, 999)
	printSlice(s3)
	printSlice(s1)

	s3 = s3[:len(s3)-1]
	s3 = s3[1:]
	printSlice(s3)
	printSlice(s1)
}
