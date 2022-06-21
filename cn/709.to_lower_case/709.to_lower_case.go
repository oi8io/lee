//给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "Hello"
//输出："hello"
//
//
// 示例 2：
//
//
//输入：s = "here"
//输出："here"
//
//
// 示例 3：
//
//
//输入：s = "LOVELY"
//输出："lovely"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 100
// s 由 ASCII 字符集中的可打印字符组成
//
// 👍 204 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func toLowerCase(s string) string {
	var r = make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 {
			r[i] = s[i] + 32
		} else {
			r[i] = s[i]
		}
	}
	return string(r)
}

//leetcode submit region end(Prohibit modification and deletion)
