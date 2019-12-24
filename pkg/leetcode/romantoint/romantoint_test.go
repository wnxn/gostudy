package romantoint

import "testing"

var rtoi  = map[string]int{
	"I":1,
	"V":5,
	"X":10,
	"L":50,
	"C":100,
	"D":500,
	"M":1000,
}

var minus = map[string]string{
	"V": "I",
	"X": "I",
	"L": "X",
	"C": "X",
	"D": "C",
	"M": "C",
}

func romanToInt(s string) int {
	res :=0
	prev := ""
	for i:= len(s)-1; i>=0; i--{
		tmp := string(s[i])
		if minus[prev] ==  tmp{
			res -= rtoi[tmp]
		}else{
			res += rtoi[tmp]
		}
		prev = tmp
	}
	return res
}

func TestRoman(t *testing.T){
	tests := []struct{
		s string
		res int
	}{
		{
			"III",
			3,
		},
		{
			"IV",
			4,
		},
		{
			"LVIII",
			58,
		},
		{
			"MCMXCIV",
			1994,
		},
	}
	for _,v:=range tests{
		res := romanToInt(v.s)
		if res != v.res{
			t.Errorf("RomanToInt(%s)=expect %d, but actually %d", v.s, v.res, res)
		}
	}
}