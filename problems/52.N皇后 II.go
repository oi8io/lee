package problems

import "fmt"

//n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
//
// Related Topics 回溯 👍 338 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

var totalNQueensNum int

func totalNQueens(n int) int {
	var (
		solvedMap = make(map[int]int)
	)
	placeQueen1(1, n, solvedMap, 0)
	return totalNQueensNum
}

func printNQueens(board map[int]int) {
	for i := 1; i <= len(board); i++ {
		x := board[i]
		var line string
		for j := 1; j <= len(board); j++ {
			char := " . "
			if j == x {
				char = " Q "
			}
			line += char
		}
		fmt.Println(line)
	}
}

func placeQueen1(k int, n int, solvedMap map[int]int, ret int) {
	var abs = func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	if k > n {
		if len(solvedMap) == n {
			printNQueens(solvedMap)
			fmt.Println()
			solvedMap = make(map[int]int)
			totalNQueensNum++
		}
		return
	}
	for i := 1; i <= n; i++ {
		j := 1
		for {
			if solvedMap[j] == i { // 同一列
				break
			}
			if abs(solvedMap[j], i) == abs(k, j) { //对角线
				break
			}
			if j == k {
				break
			}
			j++
		}
		if j == k {
			solvedMap[k] = i
			placeQueen1(k+1, n, solvedMap, ret)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
