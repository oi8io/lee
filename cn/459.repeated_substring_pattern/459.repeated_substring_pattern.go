//给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
//
//
//
// 示例 1:
//
//
//输入: s = "abab"
//输出: true
//解释: 可由子串 "ab" 重复两次构成。
//
//
// 示例 2:
//
//
//输入: s = "aba"
//输出: false
//
//
// 示例 3:
//
//
//输入: s = "abcabcabcabc"
//输出: true
//解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
//
//
//
//
// 提示：
//
//
//
//
// 1 <= s.length <= 10⁴
// s 由小写英文字母组成
//
// 👍 714 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	// 如果重复次数为偶数次，则对半一定相等
	if n%2 == 0 && s[:n/2] == s[n/2:] {
		return true
	}
	for i := 3; i <= n; {
		if i%2 == 1 && n%i == 0 { // 单次重复
			if s+s[:n/i] == s[:n/i]+s { // 前面加串 == 后面加串相等
				return true
			}
		}
		i = i + 2
	}
	return false
}

func buildString(s string, n int) string {
	s0 := s[0:n]
	sx := s0
	for {
		if len(s) > 2*len(sx) {
			sx = sx + sx
		}

	}
}

//leetcode submit region end(Prohibit modification and deletion)
