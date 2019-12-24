package validbracket

import "testing"

var mp = map[byte]byte{
	'{':'}','(':')','[':']',
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i:=range s{
		if s[i] == '{' || s[i] == '(' || s[i] == '['{
			stack = append([]byte{s[i]}, stack...)
		}else{
			if len(stack) >0 &&s[i] == mp[stack[0]]{
				stack = stack[1:]
			}else{
				return false
			}
		}
	}
	if len(stack)==0{
		return true
	}else{
		return false
	}
}

func TestIsValid(t *testing.T){
	tests := []struct{
		in string
		res bool
	}{
		//{
		//	in: "(]",
		//	res: false,
		//},
		//{
		//	in: "([)]",
		//	res: false,
		//},
		{
			in: "{[]}",
			res: true,
		},
		{
			in:"",
			res: true,
		},
	}
	for _, test := range tests{
		res := isValid(test.in)
		if res != test.res{
			t.Errorf("isValid(%s)=expect %t, but actually %t", test.in,test.res, res)
		}
	}
}
