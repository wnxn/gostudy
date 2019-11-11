package two_sum

import (
	"sort"
)

func TwoSum1(nums []int, target int) []int {
	newnums := make([]int,len(nums))
	copy(newnums, nums)
	sort.Ints(newnums)

	for i:=0;i< len(newnums)-1;i++{
		for j := i+1; j< len(newnums);j++{
			if newnums[i]+newnums[j] == target{
				var res []int
				for k :=0;k<len(nums);k++{
					if nums[k]== newnums[i] || nums[k]==newnums[j]{
						res = append(res, k)
					}
				}
				return res
			}
		}
	}

	return nil
}

func TwoSum2(nums []int, target int) []int {
	mp := make(map[int]int)
	for i,v:= range nums{
		if _,ok:=mp[target-v];ok{
			return []int{mp[target-v], i}
		}
		mp[v]=i
	}
	return nil
}