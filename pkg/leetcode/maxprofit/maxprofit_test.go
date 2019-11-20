package maxprofit

import "testing"

func maxProfit(prices []int) int {
	length := len(prices)
	res := 0
	preIndex := -1
	for i := 0; i < length-1; i++ {
		// sell
		if preIndex != -1 {
			if prices[i+1] < prices[i] {
				res += prices[i] - prices[preIndex]
				preIndex = -1
			}
		} else {
			if prices[i+1] > prices[i] {
				// buy
				preIndex = i
			}
		}
	}
	if preIndex != -1 {
		res += prices[length-1] - prices[preIndex]
	}
	return res
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		profit int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			profit: 7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			profit: 4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			profit: 0,
		},
	}
	for _, test := range tests {
		res := maxProfit(test.prices)
		if test.profit != res {
			t.Errorf("maxProfit(%v) expect %d, but actually %d", test.prices, test.profit, res)
		}
	}
}
