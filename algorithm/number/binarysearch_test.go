package number

import "testing"

func TestBinarySearch(t *testing.T) {
	testcases := []struct{
		array []int
		value int
		index int
	}{
		{[]int{1,2,3,4,5},1,0},
		{[]int{1,2,3,4,5,6},6,5},
		{[]int{1,2,3,4,4,6},4,3},
		{[]int{},6,-1},
		{[]int{1,2,3,4,5},6,-1},
		{[]int{1,2,10,15,100},15, 3},
		{[]int{1,2,10,15,100},-2, -1},
		{[]int{1,2,10,15,100},101, -1},
		{[]int{1,2,10,15,100},13, -1},
		{[]int{15},15, 0},
	}

	for _,test := range testcases{
		res := BinarySearch2(test.array,test.value)
		if res != test.index{
			t.Errorf("array %v find %d, expect %d, but result is %d",
				test.array, test.value, test.index, res)
		}
	}
}
