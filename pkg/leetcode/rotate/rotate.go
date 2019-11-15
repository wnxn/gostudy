package rotate

import "fmt"

func Rotate(nums []int, k int)  {
	length := len(nums)
	if len(nums) <=1{
		return
	}
	if length == k{
		return
	}
	max := length
	for i:=0;i<=k&&max >0;i++{
		start := i
		curIndex := i
		curValue := nums[i]
		nextIndex := (i+k)%length
		for ;nextIndex!=start;nextIndex = (curIndex+k)%length{
			tmp := nums[nextIndex]
			nums[nextIndex]=curValue
			max--
			curValue = tmp
			curIndex = nextIndex
			fmt.Println(nums)
		}
		nums[nextIndex]=curValue
		max--
		fmt.Println()
		_=curIndex
	}
}