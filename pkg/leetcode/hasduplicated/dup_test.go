package hasduplicated

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	length := len(nums)
	for i:=range nums{
		if i < length-1 &&nums[i] == nums[i+1]{
			return true
		}
	}
	return false
}


