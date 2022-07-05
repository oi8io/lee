//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// 👍 4856 👎 0

package cn

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) (ans [][]int) {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		t := 0 - nums[i]
		for ii, iii := i+1, n-1; ii < n && iii > ii; {
			v := nums[ii] + nums[iii]
			if v > t {
				iii--
			} else if v < t || (ii > i+1 && nums[ii] == nums[ii-1]) {
				ii++
			} else if v == t {
				ans = append(ans, []int{nums[i], nums[ii], nums[iii]})
				iii--
				ii++
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
