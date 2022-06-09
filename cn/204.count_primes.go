//给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
//
//
//
// 示例 1：
//
//
//输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//
//
// 示例 2：
//
//
//输入：n = 0
//输出：0
//
//
// 示例 3：
//
//
//输入：n = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= n <= 5 * 10⁶
//
// 👍 893 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	var primes []int
	check := func(i int) bool {
		for j := 0; j < len(primes); j++ {
			if primes[j]*primes[j] > i {
				break
			}
			if i%primes[j] == 0 {
				return false
			}
		}
		return true
	}
	for i := 2; i < n; i++ {
		b := check(i)
		if b {
			primes = append(primes, i)
		}
	}
	return len(primes)
}

//leetcode submit region end(Prohibit modification and deletion)
