/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

package problems

// @lc code=start
func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n%5 != 0 {
			break
		}
		n = n / 5
	}
	for {
		if n%3 != 0 {
			break
		}
		n = n / 3
	}
	for {
		if n%2 != 0 {
			break
		}
		n = n / 2
	}
	return n == 1
}

// @lc code=end
