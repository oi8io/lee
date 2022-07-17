package weekly

import (
	"sort"
)

func minOperations(nums []int, numsDivide []int) int {
	sort.Ints(numsDivide)
	must := numsDivide[0]
	for i := 1; i < len(numsDivide); i++ {
		c := gcd(numsDivide[i], numsDivide[i-1])
		if c < must {
			must = c
		}
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if must%nums[i] == 0 {
			return i
		}
	}
	return -1
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
