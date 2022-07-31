package weekly304

import (
	"sort"
)

func minimumOperations(nums []int) int {
	var ans = 0
	for {
		sort.Ints(nums)
		var nnums []int
		for i := 0; i < len(nums); i++ {
			if nums[i] > 0 {
				nnums = nums[i:]
				break
			}
		}
		nums = nnums
		if len(nums) == 0 {
			return ans
		}
		if len(nums) == 1 {
			return ans + 1
		}
		ans++
		x := nums[0]
		for i := 0; i < len(nums); i++ {
			nums[i] -= x
		}
	}
	return ans
}
