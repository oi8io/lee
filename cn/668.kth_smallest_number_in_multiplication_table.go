//几乎每一个人都用 乘法表。但是你能在乘法表中快速找到第k小的数字吗？
//
// 给定高度m 、宽度n 的一张 m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。
//
// 例 1：
//
//
//输入: m = 3, n = 3, k = 5
//输出: 3
//解释:
//乘法表:
//1	2	3
//2	4	6
//3	6	9
//
//第5小的数字是 3 (1, 2, 2, 3, 3).
//
//
// 例 2：
//
//
//输入: m = 2, n = 3, k = 6
//输出: 6
//解释:
//乘法表:
//1	2	3
//2	4	6
//
//第6小的数字是 6 (1, 2, 2, 3, 4, 6).
//
//
// 注意：
//
//
// m 和 n 的范围在 [1, 30000] 之间。
// k 的范围在 [1, m * n] 之间。
//
// 👍 194 👎 0

package cn

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func findKthNumber(m int, n int, k int) int {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%02dX%02d=%02d ", i, j, i*j)
		}
		fmt.Println()
	}
	// 假设x为第k小数字
	//两次二分法
	min, max := 1, m*n
	cnt := 0
	for {
		x := (min + max) / 2
		fmt.Println(min, max, x, cnt)
		if min >= max-1 {
			return x
		}
		for i := 1; i <= m; i++ { // 统计当前行有多少少与改数字的数
			cnt += x / i
			if cnt > k || i > x {
				break
			}
		}
		fmt.Println("222->", min, max, x, cnt)
		if cnt > k {
			cnt = 0
			max = x
		} else if cnt < k {
			if min == x {
				min++
			} else {
				min = x
			}
		} else {
			return x
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
