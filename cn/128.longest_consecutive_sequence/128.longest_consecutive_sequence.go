//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
// 👍 1293 👎 0

package cn

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	answer := 0
	current := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			current++
		} else {
			if current > answer {
				answer = current
			}
			current = 1
		}
	}
	if current > answer {
		answer = current
	}
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
