//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
//
// 👍 1973 👎 0

package cn

import (
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	//	1. 确定状态
	//	1.1 最后一枚硬币一定是coins中的任意一枚
	//	1.2 分解子问题 变成 减去最后一枚硬币后所需的最少硬币数
	//	2. 转移方程
	//	f[x] = 1+min(f(x-a1),f(x-a1),f(x-an))
	//	3. 初始条件及边界情况
	//	x=0 return 0
	//	无法组合则return -1
	//	4. 计算顺序
	//  f(0) f(1)..f(x)
	// 	1 递归
	if amount == 0 {
		return 0
	}
	var answer int
	//store := make(map[int]int)
	//answer = coinChange1(coins, amount, store)
	//return answer
	answer = coinChange2(coins, amount)
	return answer
	sort.Ints(coins)
	answer = coinChange3(coins, amount, len(coins)-1, 0, math.MaxInt)
	if answer == math.MaxInt {
		return -1
	}
	return answer
}

func minCoin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func coinChange3(coins []int, amount int, i, count, ans int) int {
	if amount == 0 {
		ans = minCoin(count, ans)
	}
	if amount < 0 {
		return ans
	}
	if i < 0 || count+amount/coins[i] >= ans {
		return ans
	}
	if amount%coins[i] == 0 {
		ans = minCoin(ans, count+amount/coins[i])
		return ans
	}
	tmp := ans
	for k := amount / coins[i]; k >= 0 && k+count < ans; k-- {
		ret := coinChange3(coins, amount-k*coins[i], i-1, count+k, ans)
		tmp = minCoin(ret, tmp)
	}
	// 每种硬币可以投递0-amount/coins[i]
	return tmp
}
func coinChange2(coins []int, amount int) int {
	store := make(map[int]int)
	store[0] = 0
	for j := 1; j <= amount; j++ {
		store[j] = math.MaxInt
		for i := 0; i < len(coins); i++ {
			if j < coins[i] {
				continue
			}
			res := store[j-coins[i]]
			if res >= 0 {
				store[j] = minCoin(res+1, store[j])
			}
		}
		if store[j] == math.MaxInt {
			store[j] = -1
		}
	}
	return store[amount]
}

func coinChange1(coins []int, amount int, minMap map[int]int) int {
	if v, ok := minMap[amount]; ok {
		return v
	}
	if amount == 0 {
		return 0
	}
	min := math.MaxInt
	for i := 0; i < len(coins); i++ {
		if amount < coins[i] {
			continue
		}
		res := coinChange1(coins, amount-coins[i], minMap)
		if res >= 0 && res+1 < min {
			min = res + 1
		}
	}
	if min == math.MaxInt {
		min = -1
	}
	minMap[amount] = min
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
