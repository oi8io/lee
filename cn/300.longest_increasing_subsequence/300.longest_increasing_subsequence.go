//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。
//
//
// 示例 1：
//
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
//
// 示例 3：
//
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
// 👍 2589 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	return lengthOfLISDP(nums)
	i := try(nums, 0, 1)
	return i
	return lengthOfLIS1(nums)
}
func lengthOfLISDP(nums []int) int {
	// 10,9,2,5,3,7,101,18
	//  2,3,7,101
	//  2,3,7,18
	//  2,5,7,18
	n := len(nums)
	var dp = make([]int, n)
	dp[0] = 1
	var ans = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
func lengthOfLIS1(nums []int) int {

	// 从后往左，
	// 把元素当成起始元素
	// 10,9,2,5,3,7,101,18
	//  2,3,7,101
	//  2,3,7,18
	//  2,5,7,18
	n := len(nums)
	if n == 1 {
		return 1
	}
	var dp = make([][]int, n)
	dp[n-1] = make([]int, 2)
	dp[n-1][0] = 0
	dp[n-1][1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		dp[i] = make([]int, 2)
		if nums[i] < nums[i+1] {
			dp[i][0] = max(dp[i+1][0], dp[i+1][1])
			dp[i][1] = dp[i+1][1] + 1
		} else {
			dp[i][0] = max(dp[i+1][0], dp[i+1][1])
			dp[i][1] = dp[i+1][1]
		}
	}
	return max(dp[0][0], dp[0][1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func try(nums []int, cur int, cnt int) int {
	if cur == len(nums)-1 {
		return cnt
	}
	m := cnt
	for i := cur + 1; i < len(nums); i++ {
		if nums[cur] < nums[i] {
			c1 := try(nums, i, cnt+1) //选中
			c2 := try(nums, i+1, cnt) //不选中
			c := max(c1, c2)
			m = max(c, m)
		}
	}
	return m
}

//leetcode submit region end(Prohibit modification and deletion)
