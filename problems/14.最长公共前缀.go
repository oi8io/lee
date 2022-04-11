/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (42.14%)
 * Likes:    2167
 * Dislikes: 0
 * Total Accepted:    794K
 * Total Submissions: 1.9M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] 仅由小写英文字母组成
 *
 *
 */
package problems

// import "fmt"

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	// 1. 先找到最短字符串长度 O(n)
	var si, n = 0, len(strs[0])
	for i, v := range strs {
		if len(v) < n {
			n = len(v)
			si = i
		}
	}
	s := strs[si]
	min, max := 0, n
	for {
		unMatch := false
		for _, v := range strs {
			if s[:n] != v[:n] {
				unMatch = true
				break
			}
		}
		if unMatch {
			max = n
		} else {
			min = n
			if max-1 <= min {
				break
			}
		}
		n = (max + min) / 2
	}

	// 2. 采用二分查找
	return s[:min]
}

// @lc code=end
