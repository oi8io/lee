//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//
//
//
// 示例 1:
//
//
//输入: s = "aba"
//输出: true
//
//
// 示例 2:
//
//
//输入: s = "abca"
//输出: true
//解释: 你可以删除c字符。
//
//
// 示例 3:
//
//
//输入: s = "abc"
//输出: false
//
//
//
// 提示:
//
//
// 1 <= s.length <= 10⁵
// s 由小写英文字母组成
//
// 👍 491 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	return isPalindrome1(s, false)
}
func isPalindrome1(s string, changed bool) bool {
	min, max := 0, len(s)-1
	for {
		if min >= max {
			return true
		}
		if s[min] == s[max] {
			min++
			max--
			continue
		}
		if changed {
			return false
		}
		changed = true
		if s[min+1] == s[max] {
			if s[min] == s[max-1] {
				return isPalindrome1(s[min+1:max+1], changed) || isPalindrome1(s[min:max], changed)
			} else {
				min++
			}
		} else {
			if s[min] == s[max-1] {
				max--
			} else {
				return false
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
