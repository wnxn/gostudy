package number

import "testing"

func TestBeautifulNumber(t *testing.T) {
	number :=[]int{3,13,15,127,1000,14919921443713777}
	rules:=[]int{2,3,2,2,999,496}
	for i,_ :=range number{
		v:=getRule(number[i])
		if v != rules[i]{
			t.Errorf("number of %d expect %d, but has %d",
				number[i], rules[i],v)
		}
	}
}
