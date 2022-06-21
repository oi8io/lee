//给你一个大小为 m x n 的网格和一个球。球的起始坐标为 [startRow, startColumn] 。你可以将球移到在四个方向上相邻的单元格内（可以
//穿过网格边界到达网格之外）。你 最多 可以移动 maxMove 次球。
//
// 给你五个整数 m、n、maxMove、startRow 以及 startColumn ，找出并返回可以将球移出边界的路径数量。因为答案可能非常大，返回对
//10⁹ + 7 取余 后的结果。
//
//
//
// 示例 1：
//
//
//输入：m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
//输出：6
//
//
// 示例 2：
//
//
//输入：m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
//输出：12
//
//
//
//
// 提示：
//
//
// 1 <= m, n <= 50
// 0 <= maxMove <= 50
// 0 <= startRow < m
// 0 <= startColumn < n
//
// 👍 246 👎 0

package cn

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	/**
	1. 确定状态 startRow+maxMove>m || startColumn+maxMove>n
	2. 转移方程 nope
	3. 边界条件
	4. 计算顺序

	*/
	var answer int
	var dp = make(map[int]int)
	a := findPaths1(m, n, maxMove, startRow, startColumn, dp, &answer)
	return a
}

// 回溯法 fixme 这个题我真做出来了吗？为什么答案是对的
func findPaths2(m int, n int, maxMove int, startRow int, startColumn int, dp map[int]int, answer *int) int {
	/**
	1. 角落有两种路径（上左，下右，上右，下右）
	2. 边上有一条路径
	3. 其余的都只能消耗路径
	*/
	idx := getIndex(maxMove, startRow, startColumn)
	if v, ok := dp[idx]; ok {
		return v
	}
	if startRow == -1 || startRow == m { // 上下已经走出框内
		*answer++
		dp[idx] += 1
		return 1
	}
	if startColumn == -1 || startColumn == n { //左右已经走出框内
		*answer++
		dp[idx] += 1
		return 1
	}
	if maxMove == 0 { // 没有步骤了，不能再走了
		dp[idx] = 0
		return 0
	}
	maxMove--
	up := dp[getIndex(maxMove, startRow-1, startColumn)]    //上
	down := dp[getIndex(maxMove, startRow+1, startColumn)]  //下
	left := dp[getIndex(maxMove, startRow, startColumn-1)]  //左
	right := dp[getIndex(maxMove, startRow, startColumn+1)] //右
	//up := findPaths1(m, n, maxMove, startRow-1, startColumn, dp, answer)    //上
	//down := findPaths1(m, n, maxMove, startRow+1, startColumn, dp, answer)  //下
	//left := findPaths1(m, n, maxMove, startRow, startColumn-1, dp, answer)  //左
	//right := findPaths1(m, n, maxMove, startRow, startColumn+1, dp, answer) //右
	var ans int
	ans = up + down + left + right
	dp[idx] = ans % (int(math.Pow(10, 9)) + 7)
	return dp[idx]
}
func getIndex(r, c, m int) int {
	idx := m*10000 + r*100 + c
	return idx
}

// 回溯法 fixme 这个题我真做出来了吗？为什么答案是对的
func findPaths1(m int, n int, maxMove int, startRow int, startColumn int, dp map[int]int, answer *int) int {
	/**
	1. 角落有两种路径（上左，下右，上右，下右）
	2. 边上有一条路径
	3. 其余的都只能消耗路径
	*/
	idx := maxMove*10000 + startRow*100 + startColumn
	if v, ok := dp[idx]; ok {
		return v
	}
	if startRow == -1 || startRow == m { // 上下已经走出框内
		*answer++
		dp[idx] += 1
		return 1
	}
	if startColumn == -1 || startColumn == n { //左右已经走出框内
		*answer++
		dp[idx] += 1
		return 1
	}
	if maxMove == 0 { // 没有步骤了，不能再走了
		dp[idx] = 0
		return 0
	}
	maxMove--
	up := findPaths1(m, n, maxMove, startRow-1, startColumn, dp, answer)    //上
	down := findPaths1(m, n, maxMove, startRow+1, startColumn, dp, answer)  //下
	left := findPaths1(m, n, maxMove, startRow, startColumn-1, dp, answer)  //左
	right := findPaths1(m, n, maxMove, startRow, startColumn+1, dp, answer) //右
	var ans int
	ans = up + down + left + right
	dp[idx] = ans % (int(math.Pow(10, 9)) + 7)
	return dp[idx]
}

//leetcode submit region end(Prohibit modification and deletion)
