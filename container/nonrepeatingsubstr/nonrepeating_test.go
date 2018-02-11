package main

import (
	"testing"
)

func TestSubarray(t *testing.T){
	testcase := []struct{
		str string
		len int
	}{
		{"asdadf", 3},
		{"一二三二一",3},
		{"pwwkew",3},
		{"这里是慕渴望",6},
	}
	for _,v:=range testcase{
		actual := maxSubarray(v.str)
		if actual != v.len{
			t.Errorf("max subarray of %s, acutal is %d, expect %d",
				v.str, actual, v.len)
		}
	}
}

func BenchmarkSubarray(b *testing.B){
	s:="这里是慕课网"
	for i := 0; i<13;i++{
		s=s+s
	}
	ans:=6
	b.Logf("len(s)=%d", len(s))
	b.ResetTimer()
	for i:=0; i<b.N;i++{
		actual:=maxSubarray(s)
		if actual!=ans{
			b.Errorf("got %d for input %s, expected %d", actual,s,ans)
		}
	}


}
