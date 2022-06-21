//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 5
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 20
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= n, m <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// 每行的所有元素从左到右升序排列
// 每列的所有元素从上到下升序排列
// -10⁹ <= target <= 10⁹
//
// 👍 1042 👎 0

package cn

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	// 对角线二分查找，以长的为准，如果较短的已经到达了最后一个index则用最后一个index遍历
	for i := 0; i < len(matrix); i++ {
		var x []string
		for j := 0; j < len(matrix[0]); j++ {
			x = append(x, strconv.Itoa(matrix[i][j]))
		}
		fmt.Println(strings.Join(x, ","))

	}
	var (
		m = len(matrix)
		n = len(matrix[0])
		i = 0
		j = n - 1
	)
	for {
		if matrix[i][j] == target {
			return true
		}
		if target < matrix[i][j] {
			if j > 0 {
				j--
			} else {
				break
			}
		} else {
			if i < m-1 {
				i++
			} else {
				break
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
