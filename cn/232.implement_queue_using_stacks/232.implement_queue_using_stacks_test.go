package cn

import (
	"testing"
)

func TestX(t *testing.T) {
	var c = Constructor()
	//["MyQueue","push","push","peek","pop","empty"]
	//	[[],[1],[2],[],[],[]]

	c.Push(1)
	c.Push(2)
	c.Peek()
	c.Pop()
	c.Empty()

}
