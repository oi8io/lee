//给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
//请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
//
// 你必须找到一个内存复杂度优于 O(n²) 的解决方案。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
//输出：13
//解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13
//
//
// 示例 2：
//
//
//输入：matrix = [[-5]], k = 1
//输出：-5
//
//
//
//
// 提示：
//
//
// n == matrix.length
// n == matrix[i].length
// 1 <= n <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// 题目数据 保证 matrix 中的所有行和列都按 非递减顺序 排列
// 1 <= k <= n²
//
//
//
//
// 进阶：
//
//
// 你能否用一个恒定的内存(即 O(1) 内存复杂度)来解决这个问题?
// 你能在 O(n) 的时间复杂度下解决这个问题吗?这个方法对于面试来说可能太超前了，但是你会发现阅读这篇文章（ this paper ）很有趣。
//
// 👍 828 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func kthSmallest(matrix [][]int, k int) int {
	return kthSmallest2(matrix, k)
	return kthSmallest1(matrix, k)
}

//二分查找
func kthSmallest2(matrix [][]int, k int) int {
	n := len(matrix)
	lo, hi := matrix[0][0], matrix[n-1][n-1] // 第一个数和最后一个数 中间
	for lo < hi {
		mi := lo + (hi-lo)>>1
		if getCount(matrix, mi, n) >= k {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo
}

// 统计所有小于mi的数字
func getCount(matrix [][]int, mi, n int) int {
	i, j := n-1, 0
	cnt := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mi {
			cnt += i + 1 // 纵向统计
			j++
		} else {
			i--
		}
	}
	return cnt
}

func kthSmallest1(matrix [][]int, k int) int {
	var ans = make([]int, len(matrix[0]))
	copy(ans, matrix[0])
	for i := 1; i < len(matrix); i++ {
		ans = merge(ans, matrix[i])
	}
	return ans[k-1]
}

func merge(a, b []int) []int {
	i, j := 0, 0
	r := make([]int, 0)
	for {
		if i == len(a) {
			r = append(r, b[j:]...)
			return r
		}
		if j == len(b) {
			r = append(r, a[i:]...)
			return r
		}
		if a[i] > b[j] {
			r = append(r, b[j])
			j++
		} else {
			r = append(r, a[i])
			i++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
