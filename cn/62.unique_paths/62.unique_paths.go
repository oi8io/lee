//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
// 问总共有多少条不同的路径？
//
//
//
// 示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//
// 示例 2：
//
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//
//
// 示例 3：
//
//
//输入：m = 7, n = 3
//输出：28
//
//
// 示例 4：
//
//
//输入：m = 3, n = 3
//输出：6
//
//
//
// 提示：
//
//
// 1 <= m, n <= 100
// 题目数据保证答案小于等于 2 * 10⁹
//
// 👍 1437 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	//	1. 确定状态，最后位置在又下角
	//	2. 到每个位置的可以分别由上或者左边的格子到达 dp[i][j] = dp[i-1][j]+dp[i][j-1]
	//	3. 初始条件，第0行第0列为0，边界条件：最右边只能向下，最下边只能向右
	//  4. 计算顺序
	var dp = make([][]int, m, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
