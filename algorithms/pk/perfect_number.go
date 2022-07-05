package pk

import "fmt"

func GetNumbers(n int) {
	for a := 2; a <= n; a++ {
		for b := 2; b < a; b++ {
			for c := b + 1; c < a; c++ {
				for d := c + 1; d < a; d++ {
					if Pow3(a) == Pow3(b)+Pow3(c)+Pow3(d) {
						fmt.Printf("a:%d=>(%d,%d,%d)\n", a, b, c, d)
					}
				}
			}
		}
	}
}

func Pow3(i int) int {
	return i * i * i
}
