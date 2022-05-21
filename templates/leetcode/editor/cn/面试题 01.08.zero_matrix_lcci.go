//编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
//
//
//
// 示例 1：
//
// 输入：
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出：
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//
//
// 示例 2：
//
// 输入：
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出：
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//
// 👍 62 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	var cols, rows []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}
	for i := 0; i < len(rows); i++ {
		matrix[rows[i]] = make([]int, len(matrix[0]))
	}
	for j := 0; j < len(cols); j++ {
		for i := 0; i < len(matrix); i++ {
			matrix[i][cols[j]] = 0
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
