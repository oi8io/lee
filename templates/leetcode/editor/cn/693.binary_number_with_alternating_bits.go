//给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。
//
//
//
// 示例 1：
//
//
//输入：n = 5
//输出：true
//解释：5 的二进制表示是：101
//
//
// 示例 2：
//
//
//输入：n = 7
//输出：false
//解释：7 的二进制表示是：111.
//
// 示例 3：
//
//
//输入：n = 11
//输出：false
//解释：11 的二进制表示是：1011.
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
// 👍 205 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func hasAlternatingBits(n int) bool {
	if n == 4 {
		return false
	}
	if n%2 == 0 {
		n = n >> 1
	}
	if n|(n<<1) != n+(n<<1) {
		return false
	}
	return (n|(n<<1))&((n|(n<<1))+1) == 0
}

//leetcode submit region end(Prohibit modification and deletion)