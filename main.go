package main

import (
	"fmt"
)

func main() {
	var x, c = 0, 0
	for i := 0; i <= 13; i++ {
		c = c ^ i
		if x != 13 {
			x = x ^ i
		}
		fmt.Println(x, c)
	}

	fmt.Println(x, c, c^x)
	fmt.Println(1 << 10)
	return
	//fmt.Println(x)
	//fmt.Println(0 ^ 1 ^ 2 ^ 3 ^ 4 ^ 5 ^ 6 ^ 7 ^ 8)
	fmt.Println(0 ^ 1 ^ 2 ^ 3 ^ 4 ^ 1)
	//ops := make([]int, 0)
	//fmt.Println(3 ^ 3)
	//fmt.Println(0 ^ 1)
	return
	ops := make([]int, 19)
	fmt.Println(ops[12])
	//ops[3:6] = []int{3, 4, 5}
	return
	ops = append(ops, 1)
	ops = append(ops, 3)
	ops = append(ops, 5)
	ops = append(ops, 7)
	ops = append(ops, 9)
	ops = ops[0:2]
	ops = append(ops, 99)
	fmt.Println(ops[:1])
	fmt.Println(ops[1:2])
	//GetSlice2(ops, 5)
	fmt.Println(ops)
	fmt.Println(ops[:1], ops[1:])
	return
	GetSlice(ops, 5)
	fmt.Println(ops)
	return
	fmt.Printf("%09b\n", 123)
	fmt.Println(9 * 9 * 9 * 9 * 9 * 9 * 9 * 9 * 9 * 9 * 9 * 9 * 9)
	var nums = []int{1, 2, 2, 3, 3, 3}
	//x := nums[1:3]
	//x = append(x, 7, 8, 10)
	//SliceInsert(nums[1:2])
	var nums1 = make([]int, len(nums))
	copy(nums1, nums)
	fmt.Println(nums1)
	nums1[3] = 88
	fmt.Println(nums)
	fmt.Println(nums1)

	return

	fmt.Printf("%c\n", byte(49))
	fmt.Println(1 << 5)
	m := make(map[int]int)
	GetMap(m)
	nw := m
	GetMap2(nw)
	fmt.Println(m)
	fmt.Println(nw)
	//fmt.Println(nums, x)
}

func GetMap(ops map[int]int) {
	ops[1] = 222
}
func GetMap2(ops map[int]int) {
	ops[5] = 222
}

func GetSlice(ops []int, op int) {
	if op == 0 {
		return
	}
	ops = append(ops, op)
	GetSlice(ops, op-1)
}
func GetSlice2(ops []int, op int) {
	ops[0] = 1
	ops = append(ops, op)
}

func GetMoney(x, y int64) {
	if y < 10 {
		fmt.Println(x, y)
		return
	}
	y = y * 20 / 100 // 20%的财富
	x = x * 98 / 100 // 98%的人
	GetMoney(x, y)
}

func SliceInsert(in []int) {
	in = append(in, 1, 2, 4)
}
