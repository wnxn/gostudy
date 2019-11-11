package myatoi

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"-13+8", -13},
		{"42", 42},
		{"     -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{" -0012a42", -12},
		{" +0 123", 0},
		{"-5-", -5},
	}
	for _, test := range tests {
		res := MyAtoi(test.input)
		if res != test.output {
			t.Errorf("MyAtoi(%s) expected %d, but actually %d", test.input, test.output, res)
		}
	}
}
