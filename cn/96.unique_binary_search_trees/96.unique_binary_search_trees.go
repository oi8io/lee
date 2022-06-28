//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：5
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 19
//
// 👍 1808 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	c := make(map[int]int)
	return tryNumTrees(n, c)
}
func tryNumTrees(n int, c map[int]int) int {
	if n <= 1 {
		return 1
	}
	if v, ok := c[n]; ok {
		return v
	}
	var answer = 0
	for i := 1; i < n+1; i++ {
		sum := tryNumTrees(i-1, c) * tryNumTrees(n-i, c)
		answer += sum
	}
	c[n] = answer
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
