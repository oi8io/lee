//请你帮忙给从 1 到 n 的数设计排列方案，使得所有的「质数」都应该被放在「质数索引」（索引从 1 开始）上；你需要返回可能的方案总数。
//
// 让我们一起来回顾一下「质数」：质数一定是大于 1 的，并且不能用两个小于它的正整数的乘积来表示。
//
// 由于答案可能会很大，所以请你返回答案 模 mod 10^9 + 7 之后的结果即可。
//
//
//
// 示例 1：
//
// 输入：n = 5
//输出：12
//解释：举个例子，[1,2,5,4,3] 是一个有效的排列，但 [5,2,3,4,1] 不是，因为在第二种情况里质数 5 被错误地放在索引为 1 的位置上。
//
//
// 示例 2：
//
// 输入：n = 100
//输出：682289015
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
//
// 👍 86 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)

const N = 1000000007

func numPrimeArrangements(n int) int {
	if n <= 2 {
		return 1
	}
	var primes = []int{2}
	for i := 3; i <= n; i++ {
		if isPrimeNum(i, primes) {
			primes = append(primes, i)
		}
	}
	n1, n2 := len(primes), n-len(primes)
	x1, x2 := getCount(n1), getCount(n2)
	// 质数全排列 * 非质数全排列
	answer := x1 * x2
	return answer % N
}
func getCount(n int) int {
	cnt := 1
	for i := 2; i <= n; i++ {
		cnt = i * cnt % N
	}
	return cnt % N
}

func isPrimeNum(n int, primes []int) bool {
	for j := 0; j < len(primes); j++ {
		if n%primes[j] == 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
