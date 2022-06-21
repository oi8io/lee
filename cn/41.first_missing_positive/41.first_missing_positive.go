//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,0]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [3,4,-1,1]
//输出：2
//
//
// 示例 3：
//
//
//输入：nums = [7,8,9,11,12]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
// 👍 1480 👎 0

package cn

import (
	"fmt"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	var cnt, total, full = 0, 0, true
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			full = false
			continue
		}
		cnt++
		total = total ^ i ^ nums[i]
		fmt.Println(total)
	}

	if full {
		return len(nums) + 1
	} else {
		return cnt ^ total
	}
}
func firstMissingPositive1(nums []int) int {
	var bin = make(map[int]uint64)
	var m uint64 = math.MaxUint64
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			continue
		}
		idx := nums[i] / 64
		bin[idx] = bin[idx] | 1<<(nums[i]%64)
	}
	if len(bin) == 0 {
		return 1
	}
	for i := 0; i < len(bin); i++ {
		if bin[i] == m {
			continue
		}
		j := 0
		if i == 0 {
			j = 1
		}
		for ; j < 64; j++ {
			if 1<<j&bin[i] == 0 {
				return i*64 + j
			}
		}
	}

	return len(bin) * 64
}

//leetcode submit region end(Prohibit modification and deletion)
