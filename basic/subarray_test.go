package main

import (
	"testing"
	"fmt"
)

func TestSubarray(t *testing.T){
	testcase := []struct{
		str string
		len int
	}{
		{"asdadf", 3},
		{"一二三二一",3},
		{"pwwkew",3},
	}
	for _,v:=range testcase{
		actual := maxSubarray(v.str)
		if actual != v.len{
			fmt.Errorf("max subarray of %s, acutal is %d, expect %d",
				v.str, actual, v.len)
		}
	}
}
