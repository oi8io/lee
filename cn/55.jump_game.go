//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// 0 <= nums[i] <= 10⁵
//
// 👍 1866 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	// 首先用回溯算法，站在当前位置尝试不同步数能否到达
	var canLast = make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] >= len(nums)-1 {
			canLast[i] = nums[i]
		}
	}
	// 如果没有一个节点能到最后节点，则返回失败
	if len(canLast) == 0 {
		return false
	}
	//fmt.Println(canLast)
	failed := make(map[int]bool)
	jump := tryCanJump(0, nums, canLast, failed)
	return jump
}

func tryCanJump(startIndex int, nums []int, canLast map[int]int, failed map[int]bool) bool {
	//fmt.Println("TRY", startIndex, nums[startIndex])
	for i := nums[startIndex]; i > 0; i-- {
		next := startIndex + i
		if next >= len(nums)-1 {
			return true
		}
		if _, ok := canLast[next]; ok {
			return true
		}
		if failed[next] {
			continue
		}
		if nums[next] == 0 {
			failed[startIndex] = true
			continue
		}
		jump := tryCanJump(next, nums, canLast, failed)
		if jump {
			return jump
		}
	}
	failed[startIndex] = true
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
