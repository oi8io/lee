//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例 1：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCCED"
//输出：true
//
//
// 示例 2：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"SEE"
//输出：true
//
//
// 示例 3：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCB"
//输出：false
//
//
//
//
// 提示：
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
//
//
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
// 👍 1356 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	exist := make(map[byte]bool)
	for i := 0; i < len(word); i++ {
		exist[word[i]] = true
	}
	pos := make(map[byte][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := board[i][j]
			if exist[c] {
				pos[c] = append(pos[c], i*n+j)
			} else {
				board[i][j] = '0'
			}
		}
	}
	visited := make(map[int]bool)
	if len(pos[word[0]]) == 0 {
		return false
	}
	b := try(board, word, 0, 0, visited, pos)
	return b
}

func try(b [][]byte, w string, i, j int, v map[int]bool, m map[byte][]int) bool {
	if i == len(w)-1 {
		return true
	}
	c := w[i]
	if len(m[c]) == 0 || j >= len(m[c]) {
		return false
	}
	v[m[w[i]][j]] = true
	x := m[c][j]
	var ret bool
	for i2, y := range m[w[i+1]] {
		if !v[y] && isLocal(len(b[0]), x, y) {
			nv := cpm(v)
			ret = try(b, w, i+1, i2, nv, m)
			if ret {
				return true
			}
		}
	}
	return try(b, w, i, j+1, v, m)
}

func cpm(m map[int]bool) map[int]bool {
	n := make(map[int]bool)
	for i, b := range m {
		n[i] = b
	}
	return n
}

func isLocal(n, x, y int) bool {
	k := abs(x, y)
	if x/n == y/n {
		return k == 1
	} else {
		return k == n
	}
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//leetcode submit region end(Prohibit modification and deletion)
