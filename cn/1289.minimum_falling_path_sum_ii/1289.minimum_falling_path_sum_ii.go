//给你一个 n x n 整数矩阵 arr ，请你返回 非零偏移下降路径 数字和的最小值。
//
// 非零偏移下降路径 定义为：从 arr 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
//
//
//
// 示例 1：
//
//
//
//
//输入：arr = [[1,2,3],[4,5,6],[7,8,9]]
//输出：13
//解释：
//所有非零偏移下降路径包括：
//[1,5,9], [1,5,7], [1,6,7], [1,6,8],
//[2,4,8], [2,4,9], [2,6,7], [2,6,8],
//[3,4,8], [3,4,9], [3,5,7], [3,5,9]
//下降路径中数字和最小的是 [1,5,7] ，所以答案是 13 。
//
//
// 示例 2：
//
//
//输入：grid = [[7]]
//输出：7
//
//
//
//
// 提示：
//
//
// n == grid.length == grid[i].length
// 1 <= n <= 200
// -99 <= grid[i][j] <= 99
//
// 👍 70 👎 0

package cn

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum(grid [][]int) int {
	/**
	1. 确定状态 最后一排取最小，每个位置分别与其不相邻的求和
	2. 转移方程：dp[i][j]=min(dp[i-1][x])+arr[i][j] x!=j && x in (0,n),由于只需要依赖上一层的结果，用两个一维数组
	3. 初始条件条件 dp[0] = grid[0] ,边界条件，如果为宽度1则为0
	4. 计算顺序，自顶向下
	*/
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	//每一行保留两个值，分别是最小是与次小值
	var dp = grid[0]
	var minIdx = 0
	for i := 0; i < len(dp); i++ {
		if dp[minIdx] > dp[i] {
			minIdx = i
		}
	}
	for i := 1; i < len(grid); i++ {
		mIdx := minIdx
		minIdx = 0
		pre := dp
		dp = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			dp[j] = math.MaxInt
			if j != mIdx {
				dp[j] = min(dp[j], pre[mIdx]+grid[i][j])
			} else {
				for x := 0; x < len(pre); x++ {
					if x != j {
						dp[j] = min(dp[j], pre[x]+grid[i][j])
					}
				}
			}
			if dp[minIdx] > dp[j] {
				minIdx = j
			}
		}
	}
	return dp[minIdx]
}

//leetcode submit region end(Prohibit modification and deletion)
