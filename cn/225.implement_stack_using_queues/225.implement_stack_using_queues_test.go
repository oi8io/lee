package cn

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	c.Push(1)
	c.Push(2)
	c.Pop()
	c.Top()
	c.Empty()
}
