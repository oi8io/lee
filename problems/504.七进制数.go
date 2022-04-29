/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 *
 * https://leetcode-cn.com/problems/base-7/description/
 *
 * algorithms
 * Easy (52.17%)
 * Likes:    177
 * Dislikes: 0
 * Total Accepted:    73.3K
 * Total Submissions: 140.6K
 * Testcase Example:  '100'
 *
 * 给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: num = 100
 * 输出: "202"
 *
 *
 * 示例 2:
 *
 *
 * 输入: num = -7
 * 输出: "-10"
 *
 *
 *
 *
 * 提示：
 *
 *
 * -10^7 <= num <= 10^7
 *
 *
 */

package problems

import "strconv"

// @lc code=start
func convertToBase7(num int) string {
	var s string
	flag := false
	if num < 0 {
		flag = true
		num = 0 - num
	}
	for {
		x := num % 7
		s = strconv.Itoa(x) + s
		num = num / 7
		if num == 0 {
			break
		}
	}
	if flag {
		s = "-" + s
	}
	return s
}

// @lc code=end
