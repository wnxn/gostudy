package game

import "testing"

func TestGame(t *testing.T) {
	tests := []struct {
		guess  []int
		answer []int
		res    int
	}{
		{
			guess:  []int{1, 2, 3},
			answer: []int{1, 2, 3},
			res:    3,
		},
		{
			guess:  []int{2, 2, 3},
			answer: []int{3, 2, 1},
			res:    1,
		},
	}
	for _, test := range tests {
		res := Game(test.guess, test.answer)
		if res != test.res {
			t.Errorf("Game(%+v, %+v)=%d, but expected %d", test.guess, test.answer, res, test.res)
		}
	}
}
