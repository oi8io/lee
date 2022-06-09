//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。
//
//
//
// 示例 1：
//
//
//输入：s = "hello"
//输出："holle"
//
//
// 示例 2：
//
//
//输入：s = "leetcode"
//输出："leotcede"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 3 * 10⁵
// s 由 可打印的 ASCII 字符组成
//
// 👍 252 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	var ret = []byte(s)
	i, j := 0, len(s)-1
	for {
		if i >= j {
			break
		}
		if !isYuan(s[i]) {
			i++
			continue
		}
		if !isYuan(s[j]) {
			j--
			continue
		}
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}
	return string(ret)
}
func isYuan(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
