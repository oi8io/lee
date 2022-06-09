//给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//
//
//
//
// 示例 1:
//
//
//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
//
// 示例 2:
//
//
//输入: numRows = 1
//输出: [[1]]
//
//
//
//
// 提示:
//
//
// 1 <= numRows <= 30
//
// 👍 761 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	var answers = make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		answers[i-1] = make([]int, i)
		answers[i-1][0] = 1
		answers[i-1][i-1] = 1
		if i <= 2 {
			continue
		}
		for j := 1; j < i-1; j++ {
			answers[i-1][j] = answers[i-2][j] + answers[i-2][j-1]
		}
	}
	return answers
}

//leetcode submit region end(Prohibit modification and deletion)
