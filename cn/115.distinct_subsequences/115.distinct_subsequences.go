//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
//
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列
//，而 "AEC" 不是）
//
// 题目数据保证答案符合 32 位带符号整数范围。
//
//
//
// 示例 1：
//
//
//输入：s = "rabbbit", t = "rabbit"
//输出：3
//解释：
//如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
//rabbbit
//rabbbit
//rabbbit
//
// 示例 2：
//
//
//输入：s = "babgbag", t = "bag"
//输出：5
//解释：
//如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。
//babgbag
//babgbag
//babgbag
//babgbag
//babgbag
//
//
//
//
// 提示：
//
//
// 0 <= s.length, t.length <= 1000
// s 和 t 由英文字母组成
//
// 👍 798 👎 0

package cn

import (
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
//babgbag
//b0 a1 , g, g
//b0 a
//第一个b,第一个a，第一个g start 位置，
func numDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	n, m := len(s), len(t)
	start, end := strings.IndexByte(s, t[0]), strings.LastIndexByte(s, t[m-1])
	if start > end || start == -1 || end == -1 {
		return 0
	}
	s = s[start : end+1]
	n = len(s)
	if n == 0 || m > n {
		return 0
	}
	var dp = make([][]int, len(s))
	dp[n-1] = make([]int, len(t))
	if s[n-1] == t[len(t)-1] {
		dp[n-1][len(t)-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		dp[i] = make([]int, len(t))
		copy(dp[i], dp[i+1])
		for j := m - 1; j >= 0; j-- {
			if s[i] == t[j] {
				if j != m-1 {
					if dp[i+1][j+1] == 0 { // 后面的字符还未出现过
						dp[i][j] = 0
					} else {
						dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
					}
				} else {
					dp[i][j]++
				}
			}
		}
	}
	return dp[0][0]
}
func numDistinct1(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	n, m := len(s), len(t)
	var exist = make(map[byte][]int)
	for i := 0; i < m; i++ {
		exist[t[i]] = append(exist[t[i]], i)
	}
	var cntMap = make([][]int, n)
	cntMap[0] = make([]int, m)
	if s[0] == t[0] {
		cntMap[0][0] = 1
	}
	for i := 1; i < n; i++ {
		cntMap[i] = make([]int, len(t))
		if i < n-1 {
			copy(cntMap[i], cntMap[i-1])
		}
		arr, ok := exist[s[i]]
		if !ok {
			continue
		}
		for x := 0; x < len(arr); x++ {
			j := arr[x]
			curBefore := cntMap[i][j]
			before := 1
			if j > 0 {
				if cntMap[i][j-1] == 0 { // 前面字母未匹配过
					before = 0
				} else {
					before = cntMap[i][j-1]
				}
			}
			cntMap[i][j] = before + curBefore
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == t[0] {
			return cntMap[i][0]
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
