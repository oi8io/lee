//给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
//
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
//
//
//
// 示例1:
//
//
//输入: pattern = "abba", str = "dog cat cat dog"
//输出: true
//
// 示例 2:
//
//
//输入:pattern = "abba", str = "dog cat cat fish"
//输出: false
//
// 示例 3:
//
//
//输入: pattern = "aaaa", str = "dog cat cat dog"
//输出: false
//
//
//
// 提示:
//
//
// 1 <= pattern.length <= 300
// pattern 只包含小写英文字母
// 1 <= s.length <= 3000
// s 只包含小写英文字母和 ' '
// s 不包含 任何前导或尾随对空格
// s 中每个单词都被 单个空格 分隔
//
// 👍 467 👎 0

package cn

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	splits := strings.Split(s, " ")
	if len(splits) != len(pattern) {
		return false
	}
	p := make(map[byte]string)
	for i, i2 := range splits {
		p[pattern[i]] = i2
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
