//给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
//
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。
//
//
//
// 示例 1：
//
//
//输入：s = "barfoothefoobarman", words = ["foo","bar"]
//输出：[0,9]
//解释：
//从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。
//
//
// 示例 2：
//
//
//输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
//输出：[]
//
//
// 示例 3：
//
//
//输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
//输出：[6,9,12]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 由小写英文字母组成
// 1 <= words.length <= 5000
// 1 <= words[i].length <= 30
// words[i] 由小写英文字母组成
//
// 👍 687 👎 0

package cn

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findSubstring(s string, words []string) []int {
	wn := len(words)
	cn := len(words[0])
	m := cn * wn
	var answer []int
	var x = make(map[string]int)
	sort.Strings(words)
	for i := 0; i < wn; i++ {
		x[words[i]] = i
	}
	// 暴力求解
	for i := 0; i <= len(s)-m; i++ {
		eq, _ := checkEq(s[i:i+m], words, x)
		if eq {
			answer = append(answer, i)
		}
	}
	return answer
}

func checkEq(s string, words []string, x map[string]int) (bool, int) {
	n := len(words[0])
	var sub []string
	for i := len(s) / n; i > 0; i-- {
		cur := s[(i-1)*n : i*n]
		if _, ok := x[cur]; ok {
			sub = append(sub, cur)
		} else {
			return false, i
		}
	}
	sort.Strings(sub)
	return isEqual(words, sub), 0
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
