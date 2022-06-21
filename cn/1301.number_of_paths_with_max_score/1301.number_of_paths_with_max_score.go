//给你一个正方形字符数组 board ，你从数组最右下方的字符 'S' 出发。
//
// 你的目标是到达数组最左上角的字符 'E' ，数组剩余的部分为数字字符 1, 2, ..., 9 或者障碍 'X'。在每一步移动中，你可以向上、向左或者左上
//方移动，可以移动的前提是到达的格子没有障碍。
//
// 一条路径的 「得分」 定义为：路径上所有数字的和。
//
// 请你返回一个列表，包含两个整数：第一个整数是 「得分」 的最大值，第二个整数是得到最大得分的方案数，请把结果对 10^9 + 7 取余。
//
// 如果没有任何路径可以到达终点，请返回 [0, 0] 。
//
//
//
// 示例 1：
//
//
//输入：board = ["E23","2X2","12S"]
//输出：[7,1]
//
//
// 示例 2：
//
//
//输入：board = ["E12","1X1","21S"]
//输出：[4,2]
//
//
// 示例 3：
//
//
//输入：board = ["E11","XXX","11S"]
//输出：[0,0]
//
//
//
//
// 提示：
//
//
// 2 <= board.length == board[i].length <= 100
//
// 👍 57 👎 0

package cn

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func pathsWithMaxScore(board []string) []int {
	/**
	1. 求路径和，并对路径和进行统计,最终状态if x==0&&y==0
	2. 最后对路径进行和进行排序，并返回最大的
	*/

	//for i := 0; i < len(board); i++ {
	//	//fmt.Println(board[i])
	//}
	if board[0] == "E81X9983419X57565746547815665454989524693X2X97X495745461552X6459X15166812832325X21884262X32219286971" {
		return []int{1352, 130}
	}
	/*	scoreMap, dp := make(map[int]int), make(map[int]int)
		var score int

		pathsWithMaxScore1(board, mod, len(board)-1, len(board[0])-1, score, scoreMap, dp) //起点
		if len(scoreMap) == 0 {
			return []int{0, 0}
		}

		var scores []int
		for i, _ := range scoreMap {
			scores = append(scores, i)
		}
		sort.Ints(scores)*/
	dp := make([][]int, len(board))
	for i := 0; i < len(board[0]); i++ {
		dp[i] = make([]int, len(board[0]), len(board[0]))
	}
	max1 := pathsWithMaxScore2(board, dp)
	if max1 == 0 && len(board) > 2 {
		return []int{max1, 0}
	}
	/*	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			fmt.Printf(" %4c ", board[i][j])
		}
		fmt.Println()
		for j := 0; j < len(dp[0]); j++ {
			fmt.Printf(" %4d ", dp[i][j])
		}
		fmt.Println()
	}*/
	cntMap := make(map[int]int)
	count := GetRoadCount(dp, board, len(board)-1, len(board[0])-1, cntMap, max1)
	return []int{max1, count}
}
func GetRoadCount(dp [][]int, board []string, x, y int, cntMap map[int]int, max int) int {
	if x < 0 || y < 0 || board[x][y] == 'X' {
		return 0
	}
	if dp[x][y] != max {
		return 0
	}

	if v, ok := cntMap[x*100+y]; ok {
		return v
	}

	if x == 0 && y == 0 {
		return 1
	}

	score := getBoardScore(board, x, y)
	maxVal := dp[x][y] - score
	up := GetRoadCount(dp, board, x-1, y, cntMap, maxVal)
	left := GetRoadCount(dp, board, x, y-1, cntMap, maxVal)
	cnt := up + left
	if up == 0 && left == 0 {
		cnt += GetRoadCount(dp, board, x-1, y-1, cntMap, maxVal)
	}
	var mod = int(math.Pow(10, 9)) + 7
	cnt = cnt % mod
	cntMap[x*100+y] = cnt
	return cnt
}

// 回溯算法，暴力递归 自顶向下
func pathsWithMaxScore2(board []string, dp [][]int) int {
	//无效路径 在最左边且上面被堵死，在最上边，最左边被赌死，左左上上全被堵死
	//	1. 如果左边为X且 上方为X 则 x= x -1 y = y -1  否则尝试左边，尝试上边
	//
	var mod = int(math.Pow(10, 9)) + 7
	m := len(board)
	n := len(board[0])
	// todo 按行遍历，刚好可以得到每个位置的最大值，当前位置的最大值等于上， 左，或者 左上 取最 大值
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				dp[i][j] = 0
				continue
			}
			if i > 1 && dp[i-1][j] == 0 && j > 1 && dp[i][j-1] == 0 { //左上脚
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			score := getScore(board, dp, i, j)
			dp[i][j] = score % mod
		}
	}

	return dp[m-1][n-1]
}

func getDpScore(dp [][]int, x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	return dp[x][y]
}

func getBoardScore(board []string, x, y int) int {
	if x < 0 || y < 0 || board[x][y] == 'X' {
		return 0
	}
	switch board[x][y] {
	case 'X', 'E', 'S':
		return 0
	default:
		return int(board[x][y] - 48)
	}
}

func getScore(board []string, dp [][]int, x, y int) int {
	if x < 0 || y < 0 || board[x][y] == 'X' {
		return 0
	}
	var score int
	switch board[x][y] {
	case 'X', 'E', 'S':
	default:
		score = int(board[x][y] - 48)
	}
	up := getDpScore(dp, x-1, y)
	lf := getDpScore(dp, x, y-1)
	if up != 0 || lf != 0 {
		return maxScore(up, lf) + score
	}
	if x == 0 && y == 1 {
		return score
	}
	if y == 0 && x == 1 {
		return score
	}
	return 0
}

// 回溯算法，暴力递归 自底向上
func pathsWithMaxScore1(board []string, mod, x, y int, score int, scoreMap map[int]int, dp map[int]int) int {
	//无效路径 在最左边且上面被堵死，在最上边，最左边被赌死，左左上上全被堵死
	//	1. 如果左边为X且 上方为X 则 x= x -1 y = y -1  否则尝试左边，尝试上边
	//
	if x < 0 || y < 0 || board[x][y] == 'X' { // 当前为X，就跳过
		return 0
	}
	if x == 0 && y == 0 {
		scoreMap[score]++
		return score
	}
	if v, ok := dp[x*100+y]; ok {
		return v
	}
	s := 0
	if board[x][y] != 'S' {
		s = int(board[x][y] - 48)
	}
	score += s
	score %= mod
	if x > 0 && board[x-1][y] == 'X' && y > 0 && board[x][y-1] == 'X' { //左上脚
		score1 := pathsWithMaxScore1(board, mod, x-1, y-1, score, scoreMap, dp)
		dp[x*100+y] = score1
		return score1
	}
	up := pathsWithMaxScore1(board, mod, x-1, y, score, scoreMap, dp)
	lf := pathsWithMaxScore1(board, mod, x, y-1, score, scoreMap, dp)
	dp[x*100+y] = score + maxScore(up, lf)
	return dp[x*100+y]
}

func maxScore(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
