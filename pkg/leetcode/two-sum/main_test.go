package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		res    []int
	}{
		{
			nums:   []int{3, 2, 3},
			target: 6,
			res:    []int{0, 2},
		},
		{
			nums:   []int{0, 4, 3, 0},
			target: 0,
			res:    []int{0, 3},
		},
	}
	for _, test := range testcases {
		res := TwoSum2(test.nums, test.target)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("TwoSum(%v, %d) expect %v, but actually %v", test.nums, test.target, test.res, res)
		}
	}
}
