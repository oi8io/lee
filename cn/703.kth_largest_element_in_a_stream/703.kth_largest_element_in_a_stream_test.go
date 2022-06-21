package cn

import (
	"fmt"
	"testing"
)

func TestConstructorKthLargest(t *testing.T) {
	constructor := Constructor(3, []int{4, 5, 8, 2})
	add := []int{3, 5, 9, 10, 4}
	for i := 0; i < len(add); i++ {
		i2 := constructor.Add(add[i])
		fmt.Println(i2)
	}
}
func TestConstructorKthLargest1(t *testing.T) {
	constructor := Constructor(1, []int{})
	add := []int{-3, -2, -4, 0, 4}
	for i := 0; i < len(add); i++ {
		i2 := constructor.Add(add[i])
		fmt.Println(i2)
	}
}
func TestConstructorKthLargest2(t *testing.T) {
	constructor := Constructor(3, []int{5, -1})
	add := []int{2, 1, -1, 3, 4}
	for i := 0; i < len(add); i++ {
		i2 := constructor.Add(add[i])
		fmt.Println(i2)
	}
}
