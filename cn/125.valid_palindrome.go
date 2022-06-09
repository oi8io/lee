//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
// 说明：本题中，我们将空字符串定义为有效的回文串。
//
//
//
// 示例 1:
//
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//解释："amanaplanacanalpanama" 是回文串
//
//
// 示例 2:
//
//
//输入: "race a car"
//输出: false
//解释："raceacar" 不是回文串
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2 * 10⁵
// 字符串 s 由 ASCII 字符组成
//
// 👍 530 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	n := len(s)
	var i, j = 0, n - 1
	for {
		if i >= j {
			return true
		}
		if !isChar(s[i]) {
			i++
			continue
		}
		if !isChar(s[j]) {
			j--
			continue
		}
		if !isPalindromeChar(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
}

func isChar(c byte) bool {
	if c >= 48 && c <= 57 { //A-Z
		return true
	}
	if c >= 65 && c <= 90 { //A-Z
		return true
	}
	if c >= 97 && c <= 122 { // a-z
		return true
	}
	return false
}
func isPalindromeChar(a, b byte) bool {
	if a > 96 {
		a = a - 32
	}
	if b > 96 {
		b = b - 32
	}
	if a == b {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
