package problems

import (
	"fmt"
)

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
// Related Topics 数组 回溯 矩阵 👍 1187 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudokuBack(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//idx := getBlockIndex(i, j)
			//val := board[i][j]
			//if board[i][j] == '.' {
			//	intVal := 1 << (int(val) - 49)
			//	for n := 1; n <= 9; n++ {
			//		if col[]
			//	}
			//
			//}
		}
	}
}

func trace(board [][]byte, x, y int) {

}

func solveSudoku(board [][]byte) {

	var (
		lines  = make([]int, 9)
		col    = make([]int, 9)
		block  = make([]int, 9)
		points = make([]int, 89)
	)

	var needFill = 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			idx := getBlockIndex(i, j)
			val := board[i][j]
			if val != '.' {
				intVal := 1 << (int(val) - 49)
				lines[i] |= intVal
				col[j] |= intVal
				block[idx] |= intVal
			} else {
				needFill++
			}
		}
	}

	// 获取可用数
	getNums := func(i, j int) int {
		nums := lines[i] | col[j] | block[getBlockIndex(i, j)]
		//fmt.Printf("(%d,%d) |%09b|%09b|%09b|%09b|%09b|%09b| \n", i, j, lines[i], col[j], block[getBlockIndex(i, j)], nums, 511^nums, 511)
		return 511 ^ nums
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val != '.' {
				continue
			}
			//找出所有可能，
			var count = 0
			nums := getNums(i, j)
			for n := 1; n <= 9; n++ {
				if 1<<n&nums == 0 {
					continue
				}
				count++
			}

			points[(i)*10+j] = nums<<4 + count
		}
	}
	//var isChanged bool

	var failed bool
	//PrintPoint(points)
	board, points, lines, col, block, needFill, _, failed = backTrack(board, points, lines, col, block, 0, needFill)

	if failed {
		fmt.Println("failed")
	}

	PrintBoard(board)

	return
}

func PrintPoint(points []int) {
	for i, bits := range points {
		count := bits & 15
		nums := bits >> 4
		if count == 0 {
			continue
		}
		fmt.Printf("（%d,%d） ---[%016b][%010b][%d] -->>>>>>\n", i/10, i%10, bits, nums, count)
	}
}

func PrintBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			fmt.Printf(" %c |", val)
		}
		fmt.Println()
	}
}

func backTrack(board [][]byte, points, lines, col, block []int, try, needFill int) (bd [][]byte, nneedFillMap, nlines, ncol, nbloc []int, ntry, nneedFill int, success bool) {
	//PrintPoint(points)
	//return board, points, lines, col, block, needFill, try, true
	if needFill == 0 {
		return board, points, lines, col, block, needFill, try, true
	}
	exist := func(i, j, n int) bool {
		if board[i][j] != '.' { //已填充
			return true
		}
		existNums := lines[i] | col[j] | block[getBlockIndex(i, j)]
		if (existNums)&1<<n != 0 { //该数字已存在
			return true
		}

		return false
	}
	place := func(i, j, n int) bool {
		num := 0 << n
		if exist(i, j, num) {
			return false
		}
		board[i][j] = byte(n + 48)
		block[getBlockIndex(i, j)] |= num
		lines[i] |= num
		col[j] |= num

		//对应的行移除可选项
		points[i*10+j] = 0
		for x := 0; x < 9; x++ {
			// 同列
			points[x*10+j] = setNum(points[x*10+j], num)
			//同行
			points[i*10+x] = setNum(points[i*10+x], num)
		}
		needFill--
		return true
	}
	cpy := func() (bd [][]byte, nneedFillMap, nlines, ncol, nblock []int, ntry, nneedFill int) {
		bd = make([][]byte, len(board))
		nlines = make([]int, 9)
		ncol = make([]int, 9)
		nblock = make([]int, 9)
		nneedFillMap = make([]int, 89)
		copy(nlines, lines)
		copy(ncol, col)
		copy(nblock, block)
		copy(nneedFillMap, points)
		copy(bd, board)
		nneedFill = needFill
		ntry = try
		return
	}
	for k := 1; k <= 9; k++ {

		var succ, isChanged bool
		for i, bits := range points {
			x, y := i/10, i%10
			count := bits & 15
			nums := bits >> 4
			if nums == 0 || count == 0 {
				continue
			}
			fmt.Printf("|---(%d,%d)[ ][%010b][%d] %v [%d]---|\n", x, y, nums, count, true, needFill)

			//fmt.Printf("Line 230 ---(%d,%d)[%d][%016b][%d] -->>>>>>\n", solveNQueensAnswers, y, solveNQueensMap, bits, count)

			xboard, xneedFillMap, xlines, xcol, xblock, xtry, xneedFill := cpy()
			//fmt.Printf("<<<<<<<<---(%d,%d)[%010b][%d]--->>>>>>\n", solveNQueensAnswers, y, nums, count)
			if k != count {
				continue
			}
			for m := 1; m <= 9; m++ {
				if isChanged && !succ { //如果放置失败，则还原后继续放置下一个可用数字
					board, points, lines, col, block, try, needFill = xboard, xneedFillMap, xlines, xcol, xblock, xtry, xneedFill
				}

				if 1<<m&nums == 0 {
					continue // 不可用
				}
				isChanged = true
				succ = place(x, y, m)
				// 优先填充唯一可以放置的数字
				if succ && try == 0 && count == 1 {
					break
				}
				fmt.Println("___+++++________+++++++")
				if !succ {
					if count == 1 {
						return board, points, lines, col, block, try, needFill, succ
					}
				} else {
					board, points, lines, col, block, needFill, try, succ = backTrack(board, points, lines, col, block, try, needFill)
					if !succ { //如果当前没有位置可选，回到上一层
						if count == 1 {
							return board, points, lines, col, block, try, needFill, succ
						}
					}
				}
				//fmt.Printf("<<<<<<<<---(%d,%d)[%d][%010b][%d] %v--->>>>>>\n", x, y, m, nums, count, succ)
			}

			if isChanged && !succ { //如果放置失败，则还原后继续放置下一个可用数字
				board, points, lines, col, block, try, needFill = xboard, xneedFillMap, xlines, xcol, xblock, xtry, xneedFill
				return board, points, lines, col, block, try, needFill, succ
			}
		}
		if isChanged {
			var faild bool
			xboard, xneedFillMap, xlines, xcol, xblock, xneedFill := board, points, lines, col, block, needFill
			board, points, lines, col, block, needFill, try, faild = backTrack(board, points, lines, col, block, try, needFill)
			if faild {
				board, points, lines, col, block, needFill = xboard, xneedFillMap, xlines, xcol, xblock, xneedFill
			}
		}
	}

	/*	for i, ints := range points {
		fmt.Println(i, ints)
	}*/
	return board, points, lines, col, block, try, needFill, true
}
func setNum(set, newNum int) int {
	if (newNum<<4)&set == 0 {
		return set
	}
	return set - (newNum<<4 + 1)
}

func getBlockIndex(x, y int) int {
	var idxMap = map[int]int{
		11: 1, 12: 2, 13: 3,
		21: 4, 22: 5, 23: 6,
		31: 7, 32: 8, 33: 9}
	var idx = ((x)/3+1)*10 + ((y)/3 + 1)
	//fmt.Println(idx)
	return idxMap[idx] - 1
}

//leetcode submit region end(Prohibit modification and deletion)
