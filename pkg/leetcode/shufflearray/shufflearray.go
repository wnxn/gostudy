package shufflearray

import (
	"math/rand"
	"time"
)

type Solution struct {
	current []int

}


func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.current
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {

	l := len(this.current)
	tmp := make([]int, l)
	copy(tmp, this.current)
	for i:=range this.current{
		index := i+rand.Int()%(l-i)
		tmp[i], tmp[index] = tmp[index], tmp[i]
	}
	return tmp
}

