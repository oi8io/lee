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
	board := make([]int, n)
	answers := make([][]int, 0)
	placeQueen(1, n, board, &answers)
	queen := printQueen(n, answers)
	return queen
}

func placeQueen(x int, n int, board []int, answers *[][]int) bool {
	if x > n {
		// 所有都摆放成功
		*answers = append(*answers, board)
		board = make([]int, n)
		return true
	}
	//fmt.Printf(answers)
	var failure int
	for y := 1; y <= n; y++ {
		if checkQueen(x, y, n, board) {
			fmt.Println(x-1, "before", board[x-1], y)
			board[x-1] = y
			sucess := placeQueen(x+1, n, board, answers)
			if sucess {
				fmt.Println(board)

			} else {
				failure++
			}
		} else {
			failure++
		}
	}
	//fmt.Printf(x, failure)
	//fmt.Printf("这是为什么啊啊啊啊啊啊")
	if failure == 4 {
		board = make([]int, 0)
		return false
	}
	return false
}

func checkQueen(x, y, n int, board []int) bool {
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	for r := 0; r < x; r++ { //遍历已经所有列
		y1 := board[r]
		if y1 == 0 {
			continue
		}
		if y == y1 { // 同一列
			fmt.Printf("+:<%d,%d><%d,%d>\n", x+1, y, r+1, y1)
			return false
		}
		if abs(x+1, r+1) == abs(y1, y) { //对角线
			fmt.Printf("x:<%d,%d><%d,%d>\n", x+1, y, r+1, y1)
			return false
		}
	}
	return true
}

func printQueen(n int, answers [][]int) [][]string {
	var ret [][]string
	for _, answer := range answers {
		var board []string
		for x := 0; x < n; x++ {
			var str string
			v := answer[x]
			fmt.Println()
			for y := 1; y <= n; y++ {
				char := "."
				if y == v {
					char = "Q"
				}
				fmt.Printf(" %s ", char)
				str += char
			}
			fmt.Println()
			board = append(board, str)
		}
		ret = append(ret, board)
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

/*
var sm *board

type board struct {
	n int
	m map[int]int
}

func newBoard(n int, m map[int]int) *board {
	return &board{n: n, m: m}
}

func (b *board) String() string {
	var strs []string
	for i := 1; i <= len(b.m); i++ {
		y := b.m[i]
		var str string
		for n := 1; n <= b.n; n++ {
			if n == y {
				str += " Q "
			} else {
				str += " + "
			}
		}
		strs = append(strs, str)
	}
	return strings.Join(strs, "\n")
}

var solveNQueensAnswers []map[int]int

func solveNQueens(n int) [][]string {
	sm = newBoard(n, make(map[int]int))
	// 第一次放第一排
	placeQueen(1, n)
	var ret [][]string
	for _, m2 := range solveNQueensAnswers {
		var board []string
		for i := 1; i <= n; i++ {
			var str string
			y := m2[i]
			for i := 1; i <= n; i++ {
				if i == y {
					str += "Q"
				} else {
					str += "."
				}
			}
			board = append(board, str)
		}
		ret = append(ret, board)
	}
	return ret
}

// 摆放第k行的皇后
func placeQueen(col int, n int) {
	if col > n { // 所有都摆放成功
		if len(sm.m) == n { //  摆放成功
			solveNQueensAnswers = append(solveNQueensAnswers, sm.m)
			sm = newBoard(n, make(map[int]int))
		}
		return
	}
	for r := 1; r <= n; r++ {
		if checkQueen(col, r, n) {
			sm.m[col] = r
			//fmt.Printf(col, r)
			//fmt.Printf(sm)
			//placeQueen(col+1, n)
		}
	}
}

func checkQueen(col, row, n int) bool {
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	for j := 1; j <= col; j++ { //遍历已经所有列
		if r, ok := sm.m[j]; ok {
			if r == row { // 同一列
				return false
			}
			if abs(r, row) == abs(col, j) { //对角线
				return false
			}
		}
	}
	return true
}
*/
