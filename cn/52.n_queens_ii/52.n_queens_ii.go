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
// 👍 360 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	numbers := n*n/64 + 1
	m := makeQueenMap(n)
	board := make([]uint64, numbers)
	//printQueenBin(m[0], n)
	return placeTotalNQueens(n, 0, board, m)
}

func placeTotalNQueens(n, x int, board []uint64, m map[int][]uint64) int {
	if x >= n {
		//printQueenBin(board, n)
		return 1
	}
	var cnt = 0
	for y := 0; y < n; y++ {
		//	纵坐标
		k := x*n + y
		isOK := checkBoard2(board, m[k])
		//isOk := checkBoard(board, n, x, y)
		//fmt.Println(n, x, y, isOk)
		if isOK {
			nb := copyBoard(board, n)
			nb[k/64] = nb[k/64] | getB(k)
			cnt += placeTotalNQueens(n, x+1, nb, m)
		}
		//	横坐标
		//	斜坐标
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
