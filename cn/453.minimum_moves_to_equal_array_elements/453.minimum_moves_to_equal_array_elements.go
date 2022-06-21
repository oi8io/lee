//给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：3
//解释：
//只需要3次操作（注意每次操作会增加两个元素的值）：
//[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1]
//输出：0
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// 答案保证符合 32-bit 整数
//
// 👍 463 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func minMoves(nums []int) int {
	//
	var min = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	var answer = 0
	// 所的数减去最小数之和
	for i := 0; i < len(nums); i++ {
		answer += nums[i] - min
	}
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
