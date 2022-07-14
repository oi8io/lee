//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
//
// 👍 1901 👎 0

package cn

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	return longestValidParentheses2(s)
}
func longestValidParentheses2(s string) int {
	var dp = make([]int, len(s))
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			j := i - dp[i-1] - 1
			if j >= 0 && s[j] == '(' {
				pre := 0
				if j > 0 {
					pre = dp[j-1]
				}
				dp[i] = dp[i-1] + 2 + pre
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
func longestValidParentheses1(s string) int {
	// 先将左边的加起来
	// 出现一个右边则左-1 有效，左为0时，为有效长度，
	//当左为负数是，则从下一个开始
	m1 := get(s)
	s1 := rev(s)
	m2 := get(s1)
	return min(m1, m2)
}
func get(s string) int {
	stack, cnt, m, bf := 0, 0, 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			cnt++
			stack++
		} else {
			stack--
			switch {
			case stack == 0:
				cnt++
				bf = cnt
			case stack > 0:
				cnt++
			case stack < 0:
				if m < cnt {
					m = cnt
				}
				cnt = 0
				bf = 0
				stack = 0
			}
		}
	}
	fmt.Println(stack)
	cnt -= stack
	if stack != 0 {
		cnt -= bf
		cnt = max(bf, cnt)
	}
	m = max(cnt, m)
	return m
}
func rev(s string) string {
	var x = make([]byte, len(s))
	for i, _ := range s {
		k := len(s) - 1 - i
		switch s[i] {
		case '(':
			x[k] = ')'
		case ')':
			x[k] = '('
		}
	}
	return string(x)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
