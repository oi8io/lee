package cn

import (
	. "github.com/oi8io/lee/cn/common"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor(GetTreeByString("1,2"))

	for i := 3; i < 18; i++ {
		c.Insert(i)
	}
	for i := 0; i < 100; i++ {
		c.Get_root()
	}
	Dump(c)
}
