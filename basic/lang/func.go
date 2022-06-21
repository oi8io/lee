package lang

import "fmt"

func nextInt(b []byte, pos int) (value, nextPos int) {
	//从i开始，直到找到数字
	for ; pos < len(b) && !isDigit(b[pos]); pos++ {
	}
	x := 0
	for ; pos < len(b) && isDigit(b[pos]); pos++ {
		x = x*10 + int(b[pos]) - '0'
	}
	return x, pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func Defer(a, b int) (c int) {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	defer func() {
		defer func() {
			fmt.Println("defer 3.1")
			c = c * 10
		}()
		fmt.Println("defer 3")
		c = c * 10
	}()

	defer func() {
		defer func() {
			fmt.Println("defer 2.1")
			c = c * 10
		}()
		fmt.Println("defer 2")
		c = c * 10
	}()
	defer func() {
		defer func() {
			fmt.Println("defer 1.1")
			c = c * 10
		}()
		fmt.Println("defer 1")
		c = c * 10
	}()
	c = a + b
	return
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
