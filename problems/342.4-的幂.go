/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

package problems

// @lc code=start
func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	for {
		if n%4 != 0 {
			break
		}
		n = n >> 2
	}
	return n == 1
}

// @lc code=end
