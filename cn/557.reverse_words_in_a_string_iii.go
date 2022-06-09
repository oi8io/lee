//给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//
//
// 示例 1：
//
//
//输入：s = "Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//
//
// 示例 2:
//
//
//输入： s = "God Ding"
//输出："doG gniD"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 5 * 10⁴
// s 包含可打印的 ASCII 字符。
// s 不包含任何开头或结尾空格。
// s 里 至少 有一个词。
// s 中的所有单词都用一个空格隔开。
//
// 👍 443 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	b := []byte(s)
	var start, end int
	for i := 0; i <= len(s); i++ {
		if i == len(s) || b[i] == ' ' {
			end = i
			reverseWord(b, start, end)
			start = i + 1
		}
	}
	return string(b)
}
func reverseWord(s []byte, start, end int) {
	for i, j := start, end-1; j >= i && i < end; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
