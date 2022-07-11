//如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
//
//
// n >= 3
// 对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
//
//
// 给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回 0 。
//
// （回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8]
//是 [3, 4, 5, 6, 7, 8] 的一个子序列）
//
//
//
//
//
//
// 示例 1：
//
//
//输入: arr = [1,2,3,4,5,6,7,8]
//输出: 5
//解释: 最长的斐波那契式子序列为 [1,2,3,5,8] 。
//
//
// 示例 2：
//
//
//输入: arr = [1,3,7,11,12,14,18]
//输出: 3
//解释: 最长的斐波那契式子序列有 [1,11,12]、[3,11,14] 以及 [7,11,18] 。
//
//
//
//
// 提示：
//
//
// 3 <= arr.length <= 1000
//
// 1 <= arr[i] < arr[i + 1] <= 10^9
//
//
// 👍 238 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func lenLongestFibSubseq(arr []int) int {
	return lenLongestFibSubseq3(arr)
	return lenLongestFibSubseq1(arr)
}
func lenLongestFibSubseq3(arr []int) int {
	n := len(arr)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[arr[i]] = i
	}
	var x = 0
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := i - 1; j >= 0 && arr[j]*2 > arr[i]; j-- {
			v := arr[i] - arr[j]
			if k, ok := m[v]; ok {
				dp[i][j] = max(dp[j][k]+1, 3)
				x = max(x, dp[i][j])
			}
		}
	}
	return x
}
func lenLongestFibSubseq1(arr []int) int {
	n := len(arr)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[arr[i]] = i
	}
	var x = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans := try(arr, m, i, j, 0)
			x = max(ans, x)
		}
	}
	return x
}
func try(arr []int, m map[int]int, i, j int, ret int) int {
	k := arr[j] + arr[i]
	if v, ok := m[k]; ok {
		if ret == 0 {
			ret = 3
		} else {
			ret++
		}
		return try(arr, m, j, v, ret)
	} else {
		return ret
	}
}
func lenLongestFibSubseq2(arr []int) int {
	// if arr[x]=arr[i] + arr[j] => dp[x]=dp[j]+1
	// 2,4,5,6,7,8,11,13,14,15,21,22,34
	// 2 4 6 x
	// 2 5 7 12 x
	// 2 6 8 14 15
	// 4 5
	// dp[6][4] = dp[4][2]+1
	n := len(arr)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[arr[i]] = i
	}
	dp := make([]int, len(arr))
	dp[0], dp[1] = 1, 1
	for i := 2; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j > 0; j-- {
			k := arr[i] - arr[j]
			if v, ok := m[k]; ok && v < j {
				dp[i] = max(dp[v]+2, dp[i])
			}
		}
	}

	return dp[n-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
