//峰值元素是指其值严格大于左右相邻值的元素。
//
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
// 你可以假设 nums[-1] = nums[n] = -∞ 。
//
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,1]
//输出：2
//解释：3 是峰值元素，你的函数应该返回其索引 2。
//
// 示例 2：
//
//
//输入：nums = [1,2,1,3,5,6,4]
//输出：1 或 5
//解释：你的函数可以返回索引 1，其峰值元素为 2；
//     或者返回索引 5， 其峰值元素为 6。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// -2³¹ <= nums[i] <= 2³¹ - 1
// 对于所有有效的 i 都有 nums[i] != nums[i + 1]
//
// 👍 846 👎 0

package cn

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findPeakElement(nums []int) (idx int) {
	return findPeakElement1(nums)
	return findPeakElement2(nums)
}
func findPeakElement2(nums []int) (idx int) {
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return idx
}
func findPeakElement1(nums []int) int {
	// 增长
	lo, hi := 0, len(nums)
	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt
		}
		return nums[i]
	}
	for {
		m := lo + (hi-lo)>>1
		if lo >= hi {
			break
		}
		if get(m-1) < get(m) && get(m) > get(m+1) {
			return m
		}
		if get(m-1) > get(m) {
			hi = m
		} else if get(m) < get(m+1) {
			lo = m
		}
	}
	return lo
}

//leetcode submit region end(Prohibit modification and deletion)
