package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
	}
	for _, test := range tests {
		res := Reverse(test.input)
		if res != test.output {
			t.Errorf("Reverse(%d) expected %d, but actually %d", test.input, test.output, res)
		}
	}
}
