//编写一个程序，通过填充空格来解决数独问题。
//
// 数独的解法需 遵循如下规则：
//
//
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//
//
// 数独部分空格内已填入了数字，空白格用 '.' 表示。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".
//",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".
//","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6
//"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[
//".",".",".",".","8",".",".","7","9"]]
//输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8
//"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],[
//"4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9",
//"6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4",
//"5","2","8","6","1","7","9"]]
//解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
//
//
//
//
//
//
// 提示：
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] 是一位数字或者 '.'
// 题目数据 保证 输入数独仅有一个解
//
//
//
//
// 👍 1252 👎 0

package problems

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	//printBoard(board)
	var (
		cols   = make([]int, 9)
		rows   = make([]int, 9)
		blocks = make([]int, 9)
	)
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == 46 {
				continue
			}
			val := 1 << (board[x][y] - 49)
			b := getBlockIndex(x, y)
			rows[x] = rows[x] | val
			cols[y] = cols[y] | val
			blocks[b] = blocks[b] | val
		}
	}
	results, _ := placeSudoku(cols, rows, blocks, board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = results[i][j]
		}
	}
	//printBoard(board)
	return
}
func printBoard(board [][]byte) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			fmt.Printf(" %c ", board[x][y])
		}
		fmt.Println()
	}
}

func placeSudoku(cols, rows, blocks []int, board [][]byte) ([][]byte, bool) {
	for i := 1; i <= 9; i++ {
		for x := 0; x < 9; x++ {
			for y := 0; y < 9; y++ {
				if board[x][y] != 46 {
					continue
				}
				b := getBlockIndex(x, y)
				var valid []int
				for n := 0; n < 9; n++ {
					if checkValidNum(1<<n, rows[x], cols[y], blocks[b]) {
						valid = append(valid, n)
					}
				}
				if len(valid) == 0 { //一个可用数字都找不到
					return nil, false
				}
				if len(valid) > i { // 从1个可用数字开始，直到找不到只有1个的位置
					continue
				}
				var failure int
				for _, n := range valid {
					val := 1 << n
					//创建快照，在副本上面操作
					nCols, nRows, nBlocks, nBoard := makeCopy(cols, rows, blocks, board)
					nBoard[x][y] = byte(n + 49)
					nRows[x] = nRows[x] | val
					nCols[y] = nCols[y] | val
					nBlocks[b] = nBlocks[b] | val
					ret, success := placeSudoku(nCols, nRows, nBlocks, nBoard)
					if success {
						cols, rows, blocks, board = nCols, nRows, nBlocks, nBoard
						return ret, success //沿路返回结果
					} else {
						failure++
					}
				}
				if failure == len(valid) { // 如果一个所有数字都失败，则退出下个数字继续
					return nil, false
				}
			}
		}
	}
	return board, true
}

func getBlockIndex(x, y int) int {
	return ((x+3)/3-1)*3 + (y+3)/3 - 1
}

func makeCopy(cols, rows, blocks []int, board [][]byte) (nCols, nRows, nBlocks []int, nBoard [][]byte) {
	nCols, nRows, nBlocks, nBoard = make([]int, 9), make([]int, 9), make([]int, 9), make([][]byte, 9)
	copy(nRows, rows)
	copy(nCols, cols)
	copy(nBlocks, blocks)
	for i, _ := range board {
		nBoard[i] = make([]byte, 9)
		copy(nBoard[i], board[i])
	}
	return
}

func setNum(val, cols, rows, blocks []int, board [][]byte) {
	return
}

func checkValidNum(val, row, col, block int) bool {
	return (row|col|block)&val == 0
}

//leetcode submit region end(Prohibit modification and deletion)
