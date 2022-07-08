package cn

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor([]int{-6, 10, 184})
	for i := 0; i < 50000; i++ {
		shuffle := c.Shuffle()
		reset := c.Reset()
		fmt.Print(shuffle, reset)
	}
}
