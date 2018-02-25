package number

import "fmt"

func Combinations(selected []int, array []int, num int){
	if num == 0{
		for _,v:=range selected{
			fmt.Print(v)
			fmt.Print(" ")
		}
		fmt.Println()
		return
	}
	if len(array)== 0{
		return
	}
	selected = append(selected, array[0])
	Combinations(selected, array[1:], num-1)
	selected = selected[:len(selected)-1]
	Combinations(selected, array[1:],num)
}
