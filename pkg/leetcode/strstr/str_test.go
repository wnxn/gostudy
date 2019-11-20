package strstr

import "testing"

func strStr(haystack string, needle string) int {
	len1 := len(haystack)
	len2 := len(needle)
	if len1 < len2 {
		return -1
	}
	for i := 0; i <= len1-len2; i++ {
		if haystack[i:i+len2] == needle[:] {
			return i
		}
	}
	return -1
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		res      int
	}{
		{
			haystack: "asd",
			needle:   "",
			res:      0,
		},
		{
			haystack: "hello",
			needle:   "ll",
			res:      2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			res:      -1,
		},
	}
	for _, test := range tests {
		res := strStr(test.haystack, test.needle)
		if res != test.res {
			t.Errorf("strStr(%s, %s)=%d, but expect %d", test.haystack, test.needle, res, test.res)
		}
	}
}
