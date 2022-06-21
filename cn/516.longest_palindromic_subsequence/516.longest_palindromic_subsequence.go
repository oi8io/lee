//给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
//
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
//
//
//
// 示例 1：
//
//
//输入：s = "bbbab"
//输出：4
//解释：一个可能的最长回文子序列为 "bbbb" 。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出：2
//解释：一个可能的最长回文子序列为 "bb" 。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由小写英文字母组成
//
// 👍 817 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindromeSubseq(s string) int {
	//找到一个点，左边的字符和右边的相同字符数量相等
	// 与最长公共子序列相同类型，从两边向中间靠拢
	// 当坐标相交的是就停止
	// dp[0][len(s)-1]=0 if s[0]== s[last] dp[0][len(s)-1]=1
	// dp[i][j]=max(dp[i+1][j],dp[i][j-1])
	// if s[i] ==s[j]  dp[i][j]=max(dp[i+1][j],dp[i][j-1],dp[i+1][j-1])

	//return longestPalindromeSubseqProcess1(s)
	return longestPalindromeSubseqProcess2(s)
}

/*func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/
func longestPalindromeSubseqProcess2(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([][]int, n+1, n+1)
	dp[0] = make([]int, n+1, n+1)
	var answer = 0
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1, n+1)
		for j := 1; j <= n-i; j++ {
			if s[i-1] == s[n-1-j] {
				dp[i][j] = dp[i-1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			answer = max(dp[i][j], answer)
		}
	}
	return answer
}
func longestPalindromeSubseqProcess1(s string) int {

	n := len(s)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	s1 := reverse(s)
	// 原串与逆序列的最长公共子序列，就是最长回文子序列
	answer := longestCommonSubsequence516(s, s1)
	return answer
}
func reverse(s string) string {
	var ret []byte
	for i := len(s) - 1; i >= 0; i-- {
		ret = append(ret, s[i])
	}
	return string(ret)
}

// 求最长公共子序列
func longestCommonSubsequence516(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1, m+1)
	dp[0] = make([]int, n+1, n+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1, n+1)
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			dp[i][j] = max(dp[i][j-1], dp[i][j])
			dp[i][j] = max(dp[i-1][j], dp[i][j])
		}
	}
	return dp[m][n]
}

func longestPalindromeSubseqProcess(s string, i, j int, dp map[int]int) int {
	n := len(s)
	if n == 0 || j < i {
		return 0
	}
	if n == 1 || j == i {
		return 1
	}
	if v, ok := dp[i*n+j]; ok {
		return v
	}
	maxLen := 0

	if s[i] == s[j] {
		maxLen = max(longestPalindromeSubseqProcess(s, i+1, j-1, dp)+2, maxLen)
	}
	maxLen = max(longestPalindromeSubseqProcess(s, i+1, j, dp), maxLen)
	maxLen = max(longestPalindromeSubseqProcess(s, i, j-1, dp), maxLen)

	dp[i*n+j] = maxLen
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
