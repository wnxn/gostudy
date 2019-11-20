package countandsay

import (
	"strconv"
	"testing"
)

func CountAndSay1(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		newRes := ""
		preIndex := 0
		for j := 0; j < len(res); j++ {
			if j+1 == len(res) || res[j+1] != res[preIndex] {
				newRes = newRes + strconv.Itoa(j+1-preIndex) + string(res[preIndex])
				preIndex = j + 1
			}
		}
		res = newRes

	}
	return res
}

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		res := CountAndSay(n - 1)
		preIndex := 0
		newRes := ""
		for i := 0; i < len(res); i++ {
			if i+1 == len(res) || res[preIndex] != res[i+1] {
				newRes += strconv.Itoa(i-preIndex+1) + string(res[preIndex])
				preIndex = i + 1
			}
		}
		return newRes
	}
}

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		n   int
		res string
	}{
		{
			n:   1,
			res: "1",
		},
		{
			n:   4,
			res: "1211",
		},
		{
			n:   6,
			res: "312211",
		},
	}
	for _, test := range tests {
		res := CountAndSay(test.n)
		if res != test.res {
			t.Errorf("expect %s, but actually %s", test.res, res)
		}
	}
}
