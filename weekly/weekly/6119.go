package weekly

import (
	"fmt"
	"sort"
)

func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	if n > threshold {
		return threshold + 1
	}
	sort.Ints(nums)
	fmt.Println(nums)
	for i := n - 1; i >= 0; i-- {
		hope := threshold / (n - i)
		if threshold%(n-i) != 0 {
			hope++
		}
		if nums[i] > hope {
			return n - i
		}
	}
	return -1
}
