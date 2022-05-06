/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 *
 * https://leetcode-cn.com/problems/self-dividing-numbers/description/
 *
 * algorithms
 * Easy (78.94%)
 * Likes:    233
 * Dislikes: 0
 * Total Accepted:    70.6K
 * Total Submissions: 89.5K
 * Testcase Example:  '1\n22'
 *
 * 自除数 是指可以被它包含的每一位数整除的数。
 *
 *
 * 例如，128 是一个 自除数 ，因为 128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。
 *
 *
 * 自除数 不允许包含 0 。
 *
 * 给定两个整数 left 和 right ，返回一个列表，列表的元素是范围 [left, right] 内所有的 自除数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：left = 1, right = 22
 * 输出：[1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
 *
 *
 * 示例 2:
 *
 *
 * 输入：left = 47, right = 85
 * 输出：[48,55,66,77]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= left <= right <= 10^4
 *
 *
 */

package problems

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	var ret []int
	for i := left; i <= right; i++ {
		if isSelfDividingNumbers(i) {
			ret = append(ret, i)
		}
	}
	return ret

}
func isSelfDividingNumbers(num int) bool {
	newNum := num
	for {
		if newNum == 0 {
			break
		}
		dlv := newNum % 10
		if dlv == 0 || num%dlv != 0 {
			return false
		}
		newNum = newNum / 10
	}
	return true
}

// @lc code=end
