package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct{
		nums []int
		k int
		res []int
	}{
		{
			nums: []int{1,2,3,4,5},
			k: 3,
			res: []int{3,4,5,1,2},
		},
		{
			nums: []int{1,2,3,4,5,6,7},
			k: 3,
			res: []int{5,6,7,1,2,3,4},
		},
		{
			nums: []int{1,2},
			k: 0,
			res: []int{1,2},
		},
	}
	for _, test:=range tests{
		input := make([]int,len(test.nums))
		copy(input, test.nums)
		Rotate(input, test.k)
		if !reflect.DeepEqual(input, test.res){
			t.Errorf("Rotate(%v, %d)=%v, but expect %v", test.nums, test.k, input, test.res)
		}
	}
}
