package problems

import (
	"fmt"
)

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
// Related Topics 数组 回溯 👍 1208 👎 0

/**
[0 0 0 0]
[0 0 0 0]
[0 0 0 0]
[0 0 0 0]


[1 0 0 0]
[0 0 1 0]
[0 0 0 0]
[0 0 0 0]

[1 0 0 0]
[0 0 1 0]
[0 0 0 0]
[0 1 0 0]

3,1

2,2
1,3
2,4




*/

// @lc code=start

var solveNQueensMap = make(map[int]int)
var solveNQueensAnswers []map[int]int

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
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
			//fmt.Println(str)
			board = append(board, str)
		}
		//fmt.Println()
		ret = append(ret, board)
	}
	return ret
}

// 摆放第k行的皇后
func placeQueen(k int, n int) {
	if k > n {
		if len(solveNQueensMap) == n {
			solveNQueensAnswers = append(solveNQueensAnswers, solveNQueensMap)
			solveNQueensMap = make(map[int]int)
		}
		return
	}
	for i := 1; i <= n; i++ {
		j := 1
		for {
			if solveNQueensMap[j] == i { // 同一列
				break
			}
			if abs(solveNQueensMap[j], i) == abs(k, j) { //对角线
				break
			}
			if j == k {
				break
			}
			j++
		}
		if j == k {
			solveNQueensMap[k] = i
			placeQueen(k+1, n)
		}
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type Tree struct {
	x, y     int
	level    int
	children []*Tree
}

func (t Tree) String() string {
	return fmt.Sprintf("level:%d,[%d,%d]", t.level, t.x, t.y)
}

func BuildTree(n int) *Tree {
	tree := &Tree{}
	for x := 0; x < n; x++ {
		subTree := &Tree{level: x}
		for y := 0; y < n; y++ {
			node := &Tree{x: x, y: y, level: x}
			subTree.children = append(subTree.children, node)
		}
		tree.children = append(tree.children, subTree)
	}
	return tree
}

func Lookup(tree *Tree) {
	fmt.Print(tree)
	for _, child := range tree.children {
		Lookup(child)
	}
	fmt.Println()
}

//leetcode submit region end(Prohibit modification and deletion)

// @lc code=end
