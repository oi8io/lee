//给定一个正整数 n，返回 连续正整数满足所有数字之和为 n 的组数 。
//
//
//
// 示例 1:
//
//
//输入: n = 5
//输出: 2
//解释: 5 = 2 + 3，共有两组连续整数([5],[2,3])求和后为 5。
//
// 示例 2:
//
//
//输入: n = 9
//输出: 3
//解释: 9 = 4 + 5 = 2 + 3 + 4
//
// 示例 3:
//
//
//输入: n = 15
//输出: 4
//解释: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
//
//
//
// 提示:
//
//
// 1 <= n <= 10⁹
//
// 👍 154 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func consecutiveNumbersSum(n int) int {
	// 求和公式为（n+m)/2*(m-n)
	// 暴力解法 为2开始，一直到n开方
	if n < 3 {
		return 1
	}

	// 两数相加，奇数，
	// 奇数只能由奇数个数子组成，除两个数字外
	var anwser = 1
	if n%2 == 1 {
		anwser = 2
	}
	i := 3
	//奇数 只能等于偶数个相加
	//偶数 只能等于奇数个相加
	for {
		if (i+1)*i/2 > n {
			break
		}
		avg1 := n / i
		if i%2 == 0 { //偶数个数字只能组成偶数
			if (2*avg1+1)*i == 2*n { //偶数
				//fmt.Println(i, avg1, avg1+1)
				anwser++
			}
		} else {
			if n%i == 0 {
				//fmt.Println(i, avg1)
				anwser++
			}
		}
		i++
	}
	return anwser
}

//leetcode submit region end(Prohibit modification and deletion)
