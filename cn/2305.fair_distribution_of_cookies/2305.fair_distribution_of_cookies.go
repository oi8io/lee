//给你一个整数数组 cookies ，其中 cookies[i] 表示在第 i 个零食包中的饼干数量。另给你一个整数 k 表示等待分发零食包的孩子数量，所有
//零食包都需要分发。在同一个零食包中的所有饼干都必须分发给同一个孩子，不能分开。
//
// 分发的 不公平程度 定义为单个孩子在分发过程中能够获得饼干的最大总数。
//
// 返回所有分发的最小不公平程度。
//
//
//
// 示例 1：
//
// 输入：cookies = [8,15,10,20,8], k = 2
//输出：31
//解释：一种最优方案是 [8,15,8] 和 [10,20] 。
//- 第 1 个孩子分到 [8,15,8] ，总计 8 + 15 + 8 = 31 块饼干。
//- 第 2 个孩子分到 [10,20] ，总计 10 + 20 = 30 块饼干。
//分发的不公平程度为 max(31,30) = 31 。
//可以证明不存在不公平程度小于 31 的分发方案。
//
//
// 示例 2：
//
// 输入：cookies = [6,1,3,2,2,4,1,2], k = 3
//输出：7
//解释：一种最优方案是 [6,1]、[3,2,2] 和 [4,1,2] 。
//- 第 1 个孩子分到 [6,1] ，总计 6 + 1 = 7 块饼干。
//- 第 2 个孩子分到 [3,2,2] ，总计 3 + 2 + 2 = 7 块饼干。
//- 第 3 个孩子分到 [4,1,2] ，总计 4 + 1 + 2 = 7 块饼干。
//分发的不公平程度为 max(7,7,7) = 7 。
//可以证明不存在不公平程度小于 7 的分发方案。
//
//
//
//
// 提示：
//
//
// 2 <= cookies.length <= 8
// 1 <= cookies[i] <= 10⁵
// 2 <= k <= cookies.length
//
// 👍 24 👎 0

package cn

import (
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func distributeCookies(cookies []int, k int) int {
	counts := make([]int, k)
	result := make(map[int]int)
	ans := tryDistributeCookies(cookies, 0, counts, result)
	return ans
}

// 回溯算法，todo 转 DP
func tryDistributeCookies(cookies []int, cur int, counts []int, result map[int]int) int {
	if cur == len(cookies) {
		return calFair(counts)
	}
	if v, ok := result[cur]; ok {
		if calFair(counts) >= v { // 如果已经超过了之前的尝试，就直接返回
			return v
		}
	}
	var answer = math.MaxInt
	for i := 0; i < len(counts); i++ { //每种摆放都尝试一下
		nCounts := copySlice(counts)
		nCounts[i] += cookies[cur]
		ret := tryDistributeCookies(cookies, cur+1, nCounts, result)
		answer = min(ret, answer)
	}
	result[cur] = answer
	return answer
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func copySlice(counts []int) []int {
	n := make([]int, len(counts))
	copy(n, counts)
	return n
}
func calFair(g []int) int {
	sort.Ints(g)
	return g[len(g)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
