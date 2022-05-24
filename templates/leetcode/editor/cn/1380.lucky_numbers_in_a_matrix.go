//给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。
//
// 幸运数 是指矩阵中满足同时下列两个条件的元素：
//
//
// 在同一行的所有元素中最小
// 在同一列的所有元素中最大
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
//输出：[15]
//解释：15 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
//
//
// 示例 2：
//
//
//输入：matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
//输出：[12]
//解释：12 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
//
//
// 示例 3：
//
//
//输入：matrix = [[7,8],[1,2]]
//输出：[7]
//解释：7是唯一的幸运数字，因为它是行中的最小值，列中的最大值。
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= n, m <= 50
// 1 <= matrix[i][j] <= 10^5
// 矩阵中的所有元素都是不同的
//
// 👍 119 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func luckyNumbers(matrix [][]int) []int {
	var rows, cols = make([]int, len(matrix)), make([]int, len(matrix[0]))
	var min, max = make([]int, len(matrix)), make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		min[i] = matrix[i][0]
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 {
				max[j] = matrix[i][j]
			}
			//fmt.Printf(" %2d ", matrix[i][j])
			if matrix[i][j] < min[i] {
				rows[i] = j
				min[i] = matrix[i][j]
			}
			if matrix[i][j] > max[j] {
				cols[j] = i
				max[j] = matrix[i][j]
			}
		}
		//fmt.Println()
	}
	//fmt.Println(max, min)
	//fmt.Println(cols)
	var answers []int
	for i, col := range rows {
		if cols[col] == i {
			answers = append(answers, matrix[i][col])
		}
	}
	//fmt.Println(min)

	return answers
}

//leetcode submit region end(Prohibit modification and deletion)
