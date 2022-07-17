package main

import "fmt"

func main() {
	var x int
	x = Case1(1)
	fmt.Println("result", x)
}

func Case1(a int) int {
	defer func() {
		a = a + 20
	}()
	a = 0x1000
	return a
}

func Case2(a int) (b int) {
	defer func() {
		b = 20
	}()
	b = a
	return b
}
