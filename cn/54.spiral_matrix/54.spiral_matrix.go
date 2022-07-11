//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
// 👍 1136 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	k, km := 0, min(m, n)
	var ans = make([]int, 0)
	for k < km/2 {
		i, j := k, k
		for ; j < n-k-1; j++ { // right
			ans = append(ans, matrix[i][j])
		}
		for ; i < m-k-1; i++ { // down
			ans = append(ans, matrix[i][j])
		}
		for ; j >= k+1; j-- { // left
			ans = append(ans, matrix[i][j])
		}
		for ; i >= k+1; i-- { // up
			ans = append(ans, matrix[i][j])
		}
		k++
	}
	if km%2 == 1 { //奇数处理
		if m < n { // 行
			ans = append(ans, matrix[km/2][k:n-k]...)
		} else {
			for i := km / 2; i < m-km/2; i++ {
				ans = append(ans, matrix[i][km/2])
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
