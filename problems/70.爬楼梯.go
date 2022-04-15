/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

package problems

// @lc code=start
func climbStairs(n int) int {
	var before, current int
	for i := 1; i <= n; i++ {
		if i == 1 {
			current = 1
			before = 1
		} else {
			now := before + current
			before = current
			current = now
		}
	}
	return current
}

// @lc code=end
