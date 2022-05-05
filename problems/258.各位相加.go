/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 *
 * https://leetcode-cn.com/problems/add-digits/description/
 *
 * algorithms
 * Easy (71.04%)
 * Likes:    510
 * Dislikes: 0
 * Total Accepted:    135K
 * Total Submissions: 190.1K
 * Testcase Example:  '38'
 *
 * 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: num = 38
 * 输出: 2
 * 解释: 各位相加的过程为：
 * 38 --> 3 + 8 --> 11
 * 11 --> 1 + 1 --> 2
 * 由于 2 是一位数，所以返回 2。
 *
 *
 * 示例 1:
 *
 *
 * 输入: num = 0
 * 输出: 0
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= num <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你可以不使用循环或者递归，在 O(1) 时间复杂度内解决这个问题吗？
 *
 */

package problems

// @lc code=start
func addDigits(num int) int {
	// 100*a + 10*b + c = 99*a + 9*b + (a+b+c)；所以对9取余即可。
	return (num-1)%9 + 1
}

// @lc code=end
