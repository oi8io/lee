/**
6099. 小于等于 K 的最长二进制子序列 显示英文描述
通过的用户数0
尝试过的用户数0
用户总通过次数0
用户总提交次数0
题目难度Medium
给你一个二进制字符串 s 和一个正整数 k 。

请你返回 s 的 最长 子序列，且该子序列对应的 二进制 数字小于等于 k 。

注意：

子序列可以有 前导 0 。
空字符串视为 0 。
子序列 是指从一个字符串中删除零个或者多个字符后，不改变顺序得到的剩余字符序列。


示例 1：

输入：s = "1001010", k = 5
输出：5
解释：s 中小于等于 5 的最长子序列是 "00010" ，对应的十进制数字是 2 。
注意 "00100" 和 "00101" 也是可行的最长子序列，十进制分别对应 4 和 5 。
最长子序列的长度为 5 ，所以返回 5 。
示例 2：

输入：s = "00101001", k = 1
输出：6
解释："000001" 是 s 中小于等于 1 的最长子序列，对应的十进制数字是 1 。
最长子序列的长度为 6 ，所以返回 6 。


提示：

1 <= s.length <= 1000
s[i] 要么是 '0' ，要么是 '1' 。
1 <= k <= 109
*/

package weekly

import (
	"math/bits"
	"strconv"
	"strings"
)

func longestSubsequence(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if m > n {
		return n
	}
	ans := m
	v, _ := strconv.ParseInt(s[n-m:], 2, 64)
	if int(v) > k { //最后的m位数字超过k
		ans--
	}
	// 把前面的0统计出来
	return ans + strings.Count(s[:n-m], "0")
}

func longestSubsequence0(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	var zeroMap = make(map[int]int)
	sum := 0
	if s[0] == '0' {
		zeroMap[0] = 1
	} else {
		sum = 1
	}

	var dp = make([]map[int]int, len(s), len(s))
	dp[0] = make(map[int]int)

	for i := 1; i < len(s); i++ {
		dp[i] = make(map[int]int)
		zeroMap[i] = zeroMap[i-1]
		if s[i] == '0' {
			zeroMap[i]++
		} else {
			sum += 1 << i
		}
	}
	if k >= sum {
		return len(s)
	}

	//fmt.Printf("s:%s,\nk:%0"+strconv.Itoa(len(s))+"b\n", s, k)
	/*		for i := len(s)-1; i >= 0; i-- {
			dp[i]=make(map[int]int)
			for i := 0; i < k; i++ {

			}
		}*/
	//fmt.Println(zeroMap)
	subsequence1 := longestSubsequence1(s, k, len(s)-1, 0, dp, zeroMap)
	//fmt.Println(dp)
	//fmt.Println(subsequence1)
	// 确认边界情况
	// dp[i][j]=
	return subsequence1
}

func longestSubsequence1(s string, k, current, cnt int, dp []map[int]int, zeroMap map[int]int) int {
	if current < 0 {
		//fmt.Println("end ", current, k, cnt)
		return cnt
	}
	//fmt.Printf("str [%s] current:%d,k:%d,cnt %d \n ", s[:current], current, k, cnt)

	if exist, ok := dp[current][k]; ok {
		return exist
	}
	if k == 0 {
		//fmt.Printf("end with k==1 current:%d,k:%d,cnt %d \n ", current, k, cnt)
		//fmt.Println("end1 ", current-1, k, cnt, zeroMap[len(s)-current])
		return cnt + zeroMap[current]
	}

	nk := k
	// 选中
	c := cnt
	if s[current] == '1' {
		nk -= 1 << cnt
		if nk < 0 {
			//fmt.Printf("end with nk==1 current:%d,k:%d,cnt %d  zc:%d \n ", current, k, cnt, zeroMap[current])
			c = cnt + zeroMap[current]
		} else {
			c = longestSubsequence1(s, nk, current-1, cnt+1, dp, zeroMap)
		}
	} else {
		c = longestSubsequence1(s, nk, current-1, cnt+1, dp, zeroMap)
	}
	nc := longestSubsequence1(s, k, current-1, cnt, dp, zeroMap)
	answer := max(c, nc)
	dp[current][k] = answer
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
