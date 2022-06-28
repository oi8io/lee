//给定字符串列表 strs ，返回其中 最长的特殊序列 。如果最长特殊序列不存在，返回 -1 。
//
// 特殊序列 定义如下：该序列为某字符串 独有的子序列（即不能是其他字符串的子序列）。
//
// s 的 子序列可以通过删去字符串 s 中的某些字符实现。
//
//
// 例如，"abc" 是 "aebdc" 的子序列，因为您可以删除"aebdc"中的下划线字符来得到 "abc" 。"aebdc"的子序列还包括
//"aebdc"、 "aeb" 和 "" (空字符串)。
//
//
//
//
// 示例 1：
//
//
//输入: strs = ["aba","cdc","eae"]
//输出: 3
//
//
// 示例 2:
//
//
//输入: strs = ["aaaa","aaa","aa"]
//输出: -1
//
//
//
//
// 提示:
//
//
// 2 <= strs.length <= 50
// 1 <= sub.length <= 10
// sub 只包含小写英文字母
//
// 👍 121 👎 0

package cn

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j]) || (len(strs[i]) == len(strs[j]) && strs[i] < strs[j])
	})
	notOK := make(map[int]bool)
	exist := make(map[string]int)
	exist[strs[0]] = 0
	for i := 1; i < len(strs); i++ {
		sub := strs[i]
		if v, ok := exist[sub]; ok {
			notOK[v] = true
			notOK[i] = true
			continue
		}
		exist[sub] = i
		for j := 0; j < i; j++ { // 短的只能是长的子串
			s := strs[j]
			if len(s) == len(sub) { //长度相等
				if s == sub {
					notOK[i] = true
					notOK[j] = true
				} else { //	两条都废弃
					if !notOK[j] {
						return len(s)
					}
				}
			} else {
				if !notOK[j] { //如果j位置有效的
					return len(s)
				}
				seq := isSubSeq(s, sub)
				if seq {
					notOK[i] = true
					break
				}
			}
		}
	}
	for i := 0; i < len(strs); i++ {
		if !notOK[i] {
			return len(strs[i])
		}
	}

	return -1
}

func isSubSeq(s string, sub string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if s[i] == sub[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return j == len(sub)
}

//leetcode submit region end(Prohibit modification and deletion)
