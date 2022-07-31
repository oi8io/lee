//给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
//
//
// 示例 1：
//
//
//输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//
//
// 示例 2：
//
//
//输入：s = "3[a2[c]]"
//输出："accaccacc"
//
//
// 示例 3：
//
//
//输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
//
//
// 示例 4：
//
//
//输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 30
//
// s 由小写英文字母、数字和方括号
// '[]' 组成
// s 保证是一个 有效 的输入。
// s 中所有整数的取值范围为
// [1, 300]
//
//
// 👍 1233 👎 0

package cn

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func decodeString(s string) string {
	type repeater struct {
		str   string
		times int
	}
	stack := make([]repeater, 0)
	//"2[abc]3[cd]ef"
	//3[a2[c]]

	started := -1
	var ans = ""
	for i, _ := range s {
		switch s[i] {
		case '[':
			atoi, _ := strconv.Atoi(s[started:i])
			r := repeater{times: atoi}
			stack = append(stack, r)
			started = -1
		case ']':
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			rs := getRepeatString(x.str, x.times)
			if len(stack) > 0 {
				stack[len(stack)-1].str += rs
			} else {
				ans += rs
			}
		default:
			if isNum(s[i]) {
				if started == -1 {
					started = i
				}
			} else {
				rs := string(s[i])
				if len(stack) > 0 {
					stack[len(stack)-1].str += rs
				} else {
					ans += rs
				}
			}
		}
	}
	return ans
}

func getRepeatString(str string, n int) string {
	var ans = ""
	for i := 0; i < n; i++ {
		ans += str
	}
	return ans
}
func isNum(x byte) bool {
	return x >= '0' && x <= '9'
}

//leetcode submit region end(Prohibit modification and deletion)
