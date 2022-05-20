//给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
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
//输入: rowIndex = 3
//输出: [1,3,3,1]
//
//
// 示例 2:
//
//
//输入: rowIndex = 0
//输出: [1]
//
//
// 示例 3:
//
//
//输入: rowIndex = 1
//输出: [1,1]
//
//
//
//
// 提示:
//
//
// 0 <= rowIndex <= 33
//
//
//
//
// 进阶：
//
// 你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
// 👍 398 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func getRow(rowIndex int) []int {
	var answers = make([]int, rowIndex+1)
	answers[0] = 1
	for i := 1; i <= rowIndex+1; i++ {
		answers[i-1] = 1
		if i <= 2 {
			continue
		}
		var bak = answers[0]
		for j := 1; j < i-1; j++ {
			cur := bak + answers[j]
			bak = answers[j]
			answers[j] = cur
		}
	}
	return answers
}

//leetcode submit region end(Prohibit modification and deletion)
