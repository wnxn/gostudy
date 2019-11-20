package longestcommonprefix_test

import (
	"sort"
	"testing"
)

func LongestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	if len(strs) == 0 {
		return ""
	}
	str1 := strs[0]
	str2 := strs[len(strs)-1]
	len1, len2 := len(str1), len(str2)
	if len1 > len2 {
		len1 = len2
	}
	i := 0
	for ; i < len1; i++ {
		if str1[i] != str2[i] {
			break
		}
	}
	return str1[:i]
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		res  string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			res:  "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			res:  "",
		},
	}
	for _, test := range tests {
		res := LongestCommonPrefix(test.strs)
		if res != test.res {
			t.Errorf("expect %s, but actually %s", test.res, res)
		}
	}
}
