//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的
//房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,2]
//输出：3
//解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3,1]
//输出：4
//解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 3：
//
//
//输入：nums = [1,2,3]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000
//
// 👍 1067 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {

	/**
	1. 刻画一个最优解的结构特征。
	2. 递归地定义最优解的值。
	3. 计算最优解的值，通常采用自底向上的方法。
	4. 利用计算出的信息构造一个最优解。
	*/
	if len(nums) == 1 {
		return nums[0]
	}
	var dp0, dp1 = make([]int, len(nums)), make([]int, len(nums))
	dp0[0], dp1[0] = nums[0], 0
	dp0[1], dp1[1] = nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		if i == len(nums)-1 {
			dp0[i], dp1[i] = dp0[i-1], max(nums[i]+dp1[i-2], dp1[i-1])
		} else {
			dp0[i], dp1[i] = max(nums[i]+dp0[i-2], dp0[i-1]), max(nums[i]+dp1[i-2], dp1[i-1])
		}
	}
	return max(dp0[len(nums)-1], dp1[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
