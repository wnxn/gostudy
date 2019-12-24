package intersect

import (
	"fmt"
	"reflect"
	"testing"
)

func intersect(nums1 []int, nums2 []int) []int {
	len1, len2:=len(nums1), len(nums2)
	if len1 > len2{
		len1,len2 = len2,len1
		nums1,nums2=nums2,nums1
	}
	mp:=make(map[int]int)
	res := make([]int,0)
	for _,v:=range nums1{
		mp[v]++
	}
	for _,v:=range nums2{
		fmt.Println(mp[v])
		if mp[v] > 0{
			res = append(res,v)
			mp[v]--
		}
	}
	return res
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		res   []int
	}{
		{
			nums1: []int{9, 3, 7},
			nums2: []int{6, 4, 1, 0, 0, 4, 4, 8, 7},
			res:   []int{7},
		},
		{
			nums1: []int{4, 7, 9, 7, 6, 7},
			nums2: []int{5, 0, 0, 6, 1, 6, 2, 2, 4},
			res:   []int{4},
		},
		{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			res:   []int{2, 2},
		},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			res:   []int{4, 9},
		},
	}
	for i, test := range tests {
		res := intersect(test.nums1, test.nums2)
		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("intersect(%d) expect %v, but actually %v", i, test.res, res)
		}
	}
}
