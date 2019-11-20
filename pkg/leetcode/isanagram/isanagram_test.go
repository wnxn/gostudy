package isanagram

import "testing"

func isAnagram(s string, t string) bool {
	lens, lent := len(s), len(t)
	if lens != lent {
		return false
	}
	recordS := make([]int, 26)
	recordT := make([]int, 26)
	for i := 0; i < lens; i++ {
		recordS[s[i]-'a']++
		recordT[t[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if recordS[i] != recordT[i] {
			return false
		}
	}
	return true

}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		res bool
	}{
		{
			s:   "anagram",
			t:   "nagaram",
			res: true,
		},
		{
			s:   "rat",
			t:   "car",
			res: false,
		},
	}
	for _, test := range tests {
		res := isAnagram(test.s, test.t)
		if res != test.res {
			t.Errorf("isAnagram(%s, %s)=%t, but expectly %t", test.s, test.t, res, test.res)
		}
	}
}
