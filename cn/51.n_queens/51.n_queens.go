//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
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
// 👍 1321 👎 0

package cn

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	return solveNQueen2(n)
}

func printQueenBin(b []uint64, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := i*n + j
			v := b[k/64] & getB(k)
			if v > 0 {
				fmt.Print(" o ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func QueenBinToString(boards [][]uint64, n int) [][]string {
	var answers [][]string
	for m := 0; m < len(boards); m++ {
		var answer = make([]string, n)
		b := boards[m]
		for i := 0; i < n; i++ {
			s := make([]byte, n)
			for j := 0; j < n; j++ {
				k := i*n + j
				v := b[k/64] & getB(k)
				s[j] = '.'
				if v > 0 {
					s[j] = 'Q'
				}
			}
			answer[i] = string(s)
		}
		answers = append(answers, answer)
	}
	return answers
}

func solveNQueen1(n int) [][]string {
	numbers := n*n/64 + 1
	board := make([]uint64, numbers)
	answers := make([][]uint64, 0)
	placeQueen1(n, 0, board, &answers)
	return QueenBinToString(answers, n)
}
func solveNQueen2(n int) [][]string {
	numbers := n*n/64 + 1
	board := make([]uint64, numbers)
	answers := make([][]uint64, 0)
	m := makeQueenMap(n)
	placeQueen3(n, 0, board, &answers, m)
	return QueenBinToString(answers, n)
}

func placeQueen3(n, x int, board []uint64, answers *[][]uint64, m map[int][]uint64) {
	if x >= n {
		*answers = append(*answers, board)
		return
	}
	for y := 0; y < n; y++ {
		k := x*n + y
		isOK := checkBoard2(board, m[k])
		if isOK {
			nb := copyBoard(board, n)
			nb[k/64] = nb[k/64] | getB(k)
			placeQueen1(n, x+1, nb, answers)
		}
	}
}

func checkBoard2(b, c []uint64) bool {
	for i := 0; i < len(b); i++ {
		if b[i]&c[i] > 0 {
			return false
		}
	}
	return true
}

func makeQueenMap(n int) map[int][]uint64 {
	fix := func(x, y int) []uint64 {
		var board = make([]uint64, n*n/64+1)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				k := i*n + j
				if i == x || j == y || absQueen(x, i) == absQueen(y, j) {
					board[k/64] = board[k/64] | getB(k)
				}
			}
		}
		return board
	}
	var x = make(map[int][]uint64)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x[i*n+j] = fix(i, j)
		}
	}
	return x
}

func placeQueen1(n, x int, board []uint64, answers *[][]uint64) {
	if x >= n {
		//printQueenBin(board, n)
		*answers = append(*answers, board)
		return
	}
	for y := 0; y < n; y++ {
		//	纵坐标
		isOk := checkBoard(board, n, x, y)
		//fmt.Println(n, x, y, isOk)
		if isOk {
			nb := copyBoard(board, n)
			k := x*n + y
			nb[k/64] = nb[k/64] | getB(k)
			placeQueen1(n, x+1, nb, answers)
		}
	}
}

func copyBoard(board []uint64, n int) []uint64 {
	nb := make([]uint64, n*n/64+1, n*n/64+1)
	copy(nb, board)
	return nb
}
func absQueen(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getB(k int) uint64 {
	return 1 << (k % 64)
}

func checkBoard(board []uint64, n, x, y int) bool {
	for i := 0; i < x; i++ {
		for j := 0; j < n; j++ {
			//	纵坐标 || 斜斜角
			if j == y || absQueen(x, i) == absQueen(j, y) { // 同列
				k := i*n + j
				//fmt.Println(i, j, k)
				if board[k/64]&getB(k) > 0 {
					return false
				}
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
