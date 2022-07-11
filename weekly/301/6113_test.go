package weekly

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	for i := 0; i <= 256; i++ {
		smallest := c.PopSmallest()
		fmt.Print(smallest, ",")
	}
	for i := 0; i < 1000/64; i++ {
		fmt.Printf("%064b\n", c.data[i])
	}
	fmt.Println()
	fmt.Println()
	//for i := 0; i <= 9; i++ {
	//	c.AddBack(i)
	//}
	for i := 60; i <= 70; i++ {
		c.AddBack(i)
	}
	c.AddBack(64)
	for i := 0; i < 1000/64; i++ {
		fmt.Printf("%064b\n", c.data[i])
	}

}
