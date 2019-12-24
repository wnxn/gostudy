package shufflearray

import "testing"

func TestShuffle(t *testing.T){
	a := []int{1,2,3}
	s := Constructor(a)
	t.Log(s.Reset())
	t.Log(s.Shuffle())
	t.Log(s.Reset())
	t.Log(s.Shuffle())
	t.Log(s.Shuffle())
}
