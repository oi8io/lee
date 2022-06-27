package cn

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	deserialize := c.deserialize("1,2,3,null,null,4,5")
	//fmt.Println(deserialize)
	c1 := Constructor()
	level := c1.serializeLevel(deserialize)
	fmt.Println(level)
}
