//给定两个字符串 s 和 t ，它们只包含小写字母。
//
// 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
//
// 请找出在 t 中被添加的字母。
//
//
//
// 示例 1：
//
//
//输入：s = "abcd", t = "abcde"
//输出："e"
//解释：'e' 是那个被添加的字母。
//
//
// 示例 2：
//
//
//输入：s = "", t = "y"
//输出："y"
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 1000
// t.length == s.length + 1
// s 和 t 只包含小写字母
//
// 👍 322 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findTheDifference(s string, t string) byte {
	var ms, ts = make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ms[s[i]]++
		ts[t[i]]++
	}
	ts[t[len(t)-1]]++
	for i := 'a'; i <= 'z'; i++ {
		char := byte(i)
		if ms[char] != ts[char] {
			return char
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
