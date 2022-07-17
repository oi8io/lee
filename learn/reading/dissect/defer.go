package dissect

import "fmt"

func Defer() {
	defer func() {
		defer func() {
			fmt.Println("ccc")
		}()
		fmt.Println("bbb")
	}()
	fmt.Println("aaa")
	//case1 := Case1(1)
	var x int
	x = Case1(1)
	fmt.Println("result", x)
	x = Case2(1)
	fmt.Println("result", x)
}

func Case1(a int) int {
	defer func() {
		a = 20
		fmt.Println(a)
	}()
	fmt.Println(a)
	a = 0
	return a
}

func Case2(a int) (b int) {
	defer func() {
		b = 20
		fmt.Println(b)
	}()
	fmt.Println(b)
	b = a
	fmt.Println(b)
	return b
}
