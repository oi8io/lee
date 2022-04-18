/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

package problems

// @lc code=start
func lengthOfLastWord(s string) int {
	var started bool
	var wordLen int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if started {
				return wordLen
			}
		} else {
			wordLen++
			started = true
		}

	}
	return wordLen
}

// @lc code=end
