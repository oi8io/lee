package link

import (
	"fmt"
	"testing"
)

func Test_insert(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i <= 10; i = i + 2 {
		number := findNumber(arr, i)
		arr = insert(arr, number, i)
	}
	for i := 1; i <= 10; i = i + 2 {
		number := findNumber(arr, i)
		arr = insert(arr, number, i)
	}
	fmt.Println(arr)
}

func Test_insert1(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println(i, i<<1, i<<1+1)
	}
}
