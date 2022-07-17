package effective

import (
	"errors"
	"fmt"
)

var errShortInput = errors.New("Short Input ")

// Loop Loop使用，可以提前结束循环
func Loop(src []int, validateOnly bool) (err error) {
	var (
		sizeOne = 1
		sizeTwo = 21
	)
	size := 1
	shift := 1

LoopX:
	for i := 0; i < 10; i++ {
		fmt.Println("I:", i)
	Loop0:
		for j := 0; j < 10; j++ {
			fmt.Println("J:", j)
		Loop1:
			for m := 0; m < 10; m++ {
				fmt.Println("M:", m)
			Loop2:
				for n := 0; n < 10; n++ {
					fmt.Println("N:", n)
					if n == 1 {
						fmt.Println("BREAK LoopX By N")
						continue Loop0
						break LoopX
					}
					if n == 1 {
						fmt.Println("BREAK Loop2 By N")
						break Loop2
					}
				}
				if m == 3 {
					fmt.Println("BREAK Loop2 By M")
					break Loop1
				}

			}
			if i == 2 {
				fmt.Println("Break Loop0 by I")
				break Loop0
			}
		}

	}

Loop:
	//
	for n := 0; n < len(src); n += size {
		switch {
		case src[n] < sizeOne:
			if validateOnly {
				break
			}
			size = 1
			update(src[n])

		case src[n] < sizeTwo:
			if n+1 >= len(src) {
				err = errShortInput
				fmt.Println("进来")
				break Loop
			}
			if validateOnly {
				break
			}
			size = 2
			update(src[n] + src[n+1]<<shift)
		}
	}
	return err
}
func update(int2 int) {

}

func CheckType() {
	var t interface{}
	t = int(123)
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T 打印任何类型的 t
	case bool:
		fmt.Printf("boolean %t\n", t) // t 是 bool 类型
	case int:
		fmt.Printf("integer %d\n", t) // t 是 int 类型
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
	}
}
