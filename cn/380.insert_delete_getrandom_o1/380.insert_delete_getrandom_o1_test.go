package cn

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	c.Insert(1)
	c.Insert(10)
	c.Insert(20)
	c.Insert(30)
	for i := 0; i < 10000; i++ {
		random := c.GetRandom()
		fmt.Print(random, " ")
	}
}
func TestConstructor1(t *testing.T) {
	//:["RandomizedSet","insert","insert","remove","insert","remove","getRandom"] [[],[0],[1],[0],[2],[1],[]]
	c := Constructor()
	c.Insert(0)
	c.Insert(1)
	c.Remove(0)
	fmt.Println(c)
	c.Insert(2)
	c.Remove(1)
	fmt.Println(c)
	for i := 0; i < 1; i++ {
		random := c.GetRandom()
		fmt.Print(random, " ")
	}
}
