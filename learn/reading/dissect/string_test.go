package dissect_test

import (
	"fmt"
	cn "github.com/oi8io/lee/cn/common"
	"testing"
)

func TestStruct(t *testing.T) {
	s := []byte("我爱您　Go 语言")
	for _, b := range string(s) {
		fmt.Printf("%#U", b)
	}
	fmt.Println()
	cn.Dump(s)
	var x = []byte{1, 2, 4, 6, 7, 88, 10}
	cn.Dump(x)
	x = append(x, 10)
	cn.Dump(x)
	x1 := make([]int, 8)
	cn.Dump(x1)
	x2 := make([]int, 8, 16)
	cn.Dump(x2)
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("[%c %x]", s[i], s[i])
	//}
	//fmt.Println()
	//for i, i2 := range s {
	//	fmt.Println(i, i2)
	//}

}
