package nil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNil(t *testing.T){
	var s []int
	t.Log(s==nil)
	assert.Nil(t, s)
	s = append(s,2)
	assert.Nil(t, s)
}
