package uniqchar

import (
	"testing"
)

// runtime 40 ms 60%
// memory 5.7mb 90%
//func firstUniqChar(s string) int {
//	mp := make(map[rune]int)
//	for _, v:=range s{
//		mp[v]++
//	}
//	for i,v:=range s{
//		if r,ok:=mp[v];ok&&r==1{
//			return i
//		}
//	}
//	return -1
//}

// runtime 8ms 96%
// memory 5.7MB 89%
//func firstUniqChar(s string)int{
//	m := make([]int, 26)
//
//	for i:=0;i<len(s);i++{
//		if m[s[i]-'a']==0{
//			m[s[i]-'a']=i+1
//		}else if m[s[i]-'a'] >0{
//			m[s[i]-'a'] = -1
//		}
//	}
//	res := -1
//	for _, v:=range m{
//		if v>0 &&(res==-1||res > v-1){
//			res = v-1
//		}
//	}
//	return res
//}

// runtime 60ms 27%
// memory 5.7mb 90%
func firstUniqChar(s string) int {
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; !ok {
			mp[s[i]] = i
		} else {
			mp[s[i]] = -1
		}
	}
	res := -1
	for _, v := range mp {
		if v != -1 && (res == -1 || v < res) {
			res = v
		}
	}
	return res
}

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{
		{
			s: "dabbcb",
			n: 0,
		},
		{
			s: "leetcode",
			n: 0,
		},
		{
			s: "loveleetcode",
			n: 2,
		},
	}
	for _, test := range tests {
		res := firstUniqChar(test.s)
		if res != test.n {
			t.Errorf("expect %d, but actually %d", test.n, res)
		}
	}
}
