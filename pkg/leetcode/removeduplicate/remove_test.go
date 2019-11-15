package removeduplicate

import (
	"reflect"
	"testing"
)

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <=1 {
		return length
	}
	pre := 0
	for i:=1;i<length;i++{
		if nums[i] == nums[pre]{
			continue
		}else{
			pre++
			nums[pre] = nums[i]
		}
	}
	return pre+1
}


func TestRemove(t *testing.T){
	tests := []struct{
		nums []int
		res []int
	}{
		{
			nums: []int{0,0,1,1,1,2,2,3,3,4},
			res: []int{0, 1, 2, 3, 4},
		},
		{
			nums: []int{1,1,2},
			res: []int{1,2},
		},
		{
			nums: []int{1,2,3},
			res: []int{1,2,3},
		},
	}
	for i, test := range tests{
		res := removeDuplicates(test.nums)
		if !reflect.DeepEqual(test.res, test.nums[:res]){
			t.Errorf("removeDuplicates(%v) expect %v, but actually %v", tests[i].nums,test.res, test.nums[:res] )
		}
		if res != len(test.res){
			t.Errorf("removeDuplicates(%v) expect %d, but actually %d",tests[i].nums,len(test.res),res)
		}
	}
}