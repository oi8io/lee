//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
//
//
//
// 示例 1：
//
//
//输入：num1 = "11", num2 = "123"
//输出："134"
//
//
// 示例 2：
//
//
//输入：num1 = "456", num2 = "77"
//输出："533"
//
//
// 示例 3：
//
//
//输入：num1 = "0", num2 = "0"
//输出："0"
//
//
//
//
//
//
// 提示：
//
//
// 1 <= num1.length, num2.length <= 10⁴
// num1 和num2 都只包含数字 0-9
// num1 和num2 都不包含任何前导零
//
// 👍 573 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	var answer = []byte(num1)
	nums := num2
	var m, n = len(num1), len(num2)
	if len(num2) > len(num1) {
		nums = num1
		n, m = m, n
		answer = []byte(num2)
	}
	var up = byte(0)
	for i := 1; i <= m; i++ {
		if i <= n {
			x := answer[m-i] + nums[n-i] + up - 48*2
			up = x / 10
			answer[m-i] = 48 + x%10
		} else {
			if up == 0 {
				break
			}
			x := answer[m-i] + up - 48
			up = x / 10
			answer[m-i] = 48 + x%10
		}
	}
	if up == 1 {
		return "1" + string(answer)
	}
	return string(answer)
}

//leetcode submit region end(Prohibit modification and deletion)
