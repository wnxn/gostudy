package three_sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testcases := []struct{
		nums []int
		res [][]int
	}{
		//{
		//	nums: []int{-1,0,1,2,-1,-4},
		//	res: [][]int{{-1,0,1},{-1,-1,2}},
		//},
		//{
		//	nums: []int{0,0,0,0},
		//	res: [][]int{{0,0,0}},
		//},
		{
			nums: []int{3,0,-2,-1,1,2},
			res: [][]int{
				{-2,-1,3},
				{-2,0,2},
			},
		},
	}
	for _,test:=range testcases{
		res := ThreeSum(test.nums)
		if !reflect.DeepEqual(res, test.res){
			t.Errorf("TwoSum(%v) expect %v, but actually %v", test.nums, test.res, res)
		}
	}
	
}
