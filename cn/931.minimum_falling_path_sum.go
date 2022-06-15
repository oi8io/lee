//给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
//
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第
//一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1
//, col + 1) 。
//
//
//
// 示例 1：
//
//
//
//
//输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
//输出：13
//解释：如图所示，为和最小的两条下降路径
//
//
// 示例 2：
//
//
//
//
//输入：matrix = [[-19,57],[-40,-5]]
//输出：-59
//解释：如图所示，为和最小的下降路径
//
//
//
//
// 提示：
//
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 100
// -100 <= matrix[i][j] <= 100
//
// 👍 177 👎 0

package cn

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum1(matrix [][]int) int {
	// 改题与120非常类似
	/**
	1. 确定状态，最后一行求MIN
	2. 状态转移，dp[i][j]= arr[i][j]+min(dp[i-1][j-1],dp[i-1][j],dp[i-1][j+1]))
	3. 初始条件，第一行等于其本身，边界条件 第一列没有左边元素，最后一列没有后面的元素
	4. 计算顺序，自顶而下，自下而上
	5. 复杂度
	*/
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	cur := matrix[0]
	for i := 1; i < len(matrix); i++ {
		pre := cur
		cur = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				cur[j] = matrix[i][j] + min(pre[j], pre[j+1])
			} else if j == len(matrix[i])-1 {
				cur[j] = matrix[i][j] + min(pre[j-1], pre[j])
			} else {
				cur[j] = matrix[i][j] + triMin(pre[j-1], pre[j], pre[j+1])
			}
		}
	}
	var answer = math.MaxInt
	for i := 0; i < len(cur); i++ {
		answer = min(answer, cur[i])
	}
	return answer

}
func triMin(a, b, c int) int {
	if a < b {
		if c < a {
			return c
		} else {
			return a
		}
	}
	if c < b {
		return c
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
