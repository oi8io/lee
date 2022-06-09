package cn

import (
	"fmt"
	"testing"
)

func TestMinStack_Push(t *testing.T) {
	m := ConstructorMinStack()
	m.Push(2)
	m.Push(-4)
	m.Pop()
	fmt.Println("TOP", m.Top())
	fmt.Println("MIN", m.GetMin())
	m.Push(4)
	m.Push(4)
	m.Push(4)
	m.Push(4)
	m.Push(4)
	m.Push(1)
	m.Push(1)
	m.Push(-1)
	m.Push(1)
	fmt.Println(m.Top())
	fmt.Println(m.GetMin())
}

func TestMinStack_Push1(t *testing.T) {
	//["MinStack","push","push","push","push","pop","getMin","pop","getMin","pop","getMin"]
	//	[[],[512],[-1024],[-1024],[512],[],[],[],[],[],[]]
	m := ConstructorMinStack()
	m.Push(512)
	m.Push(-1024)
	m.Push(-1024)
	m.Push(512)
	m.Pop()
	fmt.Println("TOP", m.Top())
	fmt.Println("Min", m.GetMin())
	fmt.Println("Min", m.index)
	m.Pop()
	fmt.Println()
	fmt.Println("TOP", m.Top())
	fmt.Println("Min", m.GetMin())
	fmt.Println("Min", m.index)
	m.Pop()
	fmt.Println()
	fmt.Println("TOP", m.Top())
	fmt.Println("Min", m.GetMin())
	fmt.Println("Min", m.index)
}

//Wrong Answer: input:["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"] [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]] Output:[null,null,null,null,2147483647,null,2147483646,null,2147483646,null,null,2147483647,2147483646,null,-2147483648,-2147483648,null,2147483647] Expected:[null,null,null,null,2147483647,null,2147483646,null,2147483646,null,null,2147483647,2147483647,null,-2147483648,-2147483648,null,2147483647] stdout:
