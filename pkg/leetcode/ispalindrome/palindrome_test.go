package ispalindrome

import (
	"bytes"
	"testing"
)

func isPalindrome(s string) bool {
	length := len(s)
	s = string(bytes.ToLower([]byte(s)))
	for i, j := 0, length-1; i < j; {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9')) {
			i++
			continue
		}
		if !((s[j] >= 'a' && s[j] <= 'z') || (s[j] >= '0' && s[j] <= '9')) {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s   string
		res bool
	}{
		{
			s:   "A man, a plan, a canal: Panama",
			res: true,
		},
		{
			s:   "race a car",
			res: false,
		},
		{
			s:   "0P",
			res: false,
		},
	}
	for _, test := range tests {
		res := isPalindrome(test.s)
		if res != test.res {
			t.Errorf("IsPalindrome(%s)= %t, but expect %t", test.s, res, test.res)
		}
	}
}
