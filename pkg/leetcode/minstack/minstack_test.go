package minstack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	s := MinStack{}
	s.Push(2)
	s.Push(4)
	s.Push(1)
	s.Push(2)
	assert.Equal(t, s.Top(), 2)
	assert.Equal(t, s.GetMin(), 1)
	s.Pop()
	assert.Equal(t, s.GetMin(), 1)
	s.Pop()
	assert.Equal(t, s.GetMin(), 2)
}