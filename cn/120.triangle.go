//给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果
//正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
//
//
// 示例 1：
//
//
//输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//输出：11
//解释：如下面简图所示：
//   2
//  3 4
// 6 5 7
//4 1 8 3
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//
// 示例 2：
//
//
//输入：triangle = [[-10]]
//输出：-10
//
//
//
//
// 提示：
//
//
// 1 <= triangle.length <= 200
// triangle[0].length == 1
// triangle[i].length == triangle[i - 1].length + 1
// -10⁴ <= triangle[i][j] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
//
// 👍 1051 👎 0

package cn

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	/**
	1. 确定状态 从第一行到最后一行 min (dp[i][0],... dp[0][n])
	2. 确定状态 dp[i][j] = arr[i][j]+min(dp[i-1][j],dp[i-1][j-1])
	3. 初始条件状态，边界条件
	4. 计算顺序
	*/
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	cur := triangle[0]
	for i := 1; i < len(triangle); i++ {
		dp := cur
		cur = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				cur[j] = triangle[i][j] + dp[j]
			} else {
				if j == len(triangle[i])-1 {
					cur[j] = triangle[i][j] + dp[j-1]
				} else {
					cur[j] = triangle[i][j] + min(dp[j], dp[j-1])
				}
			}
		}
	}
	answer := math.MaxInt
	for j := 0; j < len(cur); j++ {
		answer = min(cur[j], answer)
	}
	return answer

}
func minimumTotal1(triangle [][]int) int {
	/**
	1. 确定状态 从第一行到最后一行 min (dp[i][0],... dp[0][n])
	2. 确定状态 dp[i][j] = arr[i][j]+min(dp[i-1][j],dp[i-1][j-1])
	3. 初始条件状态，边界条件
	4. 计算顺序
	*/
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	var dp = make([][]int, len(triangle), len(triangle))
	dp[0] = triangle[0]
	for i := 1; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			} else {
				if j == len(triangle[i])-1 {
					dp[i][j] = triangle[i][j] + dp[i-1][j-1]
				} else {
					dp[i][j] = triangle[i][j] + min(dp[i-1][j], dp[i-1][j-1])
				}
			}
		}
	}
	answer := math.MaxInt
	arr := dp[len(triangle)-1]

	for j := 0; j < len(arr); j++ {
		answer = min(arr[j], answer)
	}

	return answer

}

//leetcode submit region end(Prohibit modification and deletion)
