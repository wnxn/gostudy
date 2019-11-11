package three_sum

import (
	"fmt"
	"sort"
)

func ThreeSum1(nums []int) [][]int {
	sort.Ints(nums)
	mp := make(map[int]int)
	for i,v:=range nums{
		mp[v]=i
	}
	var res [][]int
	fmt.Println(nums)
	var prei *int = nil
	var prej *int = nil
	for i,v:=range nums{
		if v > 0{
			return res
		}
		for j:=i+1;j<len(nums);j++{
			if k,ok:= mp[0-nums[j]-nums[i]];ok&&k>j{
				if prei !=nil && *prei == nums[i] || prej != nil && *prej == nums[j] {
					continue
				}
				r :=[]int{nums[i],nums[j],-1*(nums[j]+nums[i])}
				fmt.Println(i,j,k)
				res = append(res, r)
			}
			prej = &nums[j]
		}
		prei = &nums[i]
	}
	return res
}

func ThreeSum(nums []int) [][]int{
	sort.Ints(nums)
	var res [][]int
	if len(nums) <3{
		return nil
	}
	for i:=0;i<len(nums)-2;i++{
		if nums[i] >0{
			break
		}
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		L,R:=i+1, len(nums)-1
		for L<R && R < len(nums){
			tmp := nums[i]+ nums[L] + nums[R]
			if tmp ==0{
				res = append(res, []int{nums[i], nums[L] ,nums[R]})
				for L<R && nums[L]==nums[L+1]{
					L++
				}
				for L<R && nums[R]==nums[R-1]{
					R--
				}
				L++
				R--
			}else if tmp <0{
				L++
			}else if tmp >0{
				R--
			}
			fmt.Println(L,R,nums,i)
		}
	}
	return res
}