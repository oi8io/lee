//给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 假设你总是可以到达数组的最后一个位置。
//
//
//
// 示例 1:
//
//
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 示例 2:
//
//
//输入: nums = [2,3,0,1,4]
//输出: 2
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// 0 <= nums[i] <= 1000
//
// 👍 1660 👎 0

package cn

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
//贪心
func jump(nums []int) int {
	var pos, end, step int
	for i := 0; i < len(nums)-1; i++ {
		pos = max(pos, nums[i]+i)
		if end == i {
			end = pos
			step++
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func jump2(nums []int) int {
	// 0 到 n的 步数 = 0到x+x到n的和
	// dp[i]=min(dp[i+1]....dp[i+x]+1
	//var dp = make([][]int, len(nums))
	var dp = make([]int, len(nums))

	for i := len(nums) - 2; i >= 0; i-- { //目标位置
		min := math.MaxInt
		for j := nums[i]; j >= 1; j-- {
			cur := math.MaxInt
			if i+j >= len(nums)-1 {
				min = 1
				break
			} else {
				if dp[i+j] == 1 {
					min = 2
					break
				}
				if dp[i+j] != math.MaxInt {
					cur = dp[i+j] + 1
				}
			}
			if cur < min {
				min = cur
			}
		}
		dp[i] = min
	}
	if dp[0] == math.MaxInt {
		return 0
	}
	return dp[0]
}

func jump1(nums []int) int {
	var dp = make(map[int]int)
	i := tryJump(nums, 0, dp)
	if i == math.MaxInt {
		return 0
	}
	return i
}

func tryJump(nums []int, cur int, dp map[int]int) int {
	if cur >= len(nums)-1 {
		return 0
	}
	if exist, ok := dp[cur]; ok {
		if exist == math.MaxInt {

		}
		return exist
	}
	min := math.MaxInt
	for i := nums[cur]; i >= 1; i-- {
		step := tryJump(nums, cur+i, dp)
		if step != math.MaxInt && step+1 < min {
			min = step + 1
		}
	}
	dp[cur] = min
	return dp[cur]
}

//leetcode submit region end(Prohibit modification and deletion)
