//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// 👍 2725 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	//	类似优先级题目 全排列问题
	//	第一个括号有n种选择
	//  第二种有n-1种选择
	//fixme 每加入一个括号都有两种选择，要不放栈中
	stack := 1
	answer := try(n, 1, stack, "(")
	return answer
}
func try(n int, k int, stack int, s string) []string {
	var answer = make([]string, 0)
	if k == n {
		last := make([]byte, stack)
		for i := 0; i < stack; i++ {
			last[i] = ')'
		}
		s = s + string(last)
		return []string{s}
	}
	a := try(n, k+1, stack+1, s+"(")
	answer = append(answer, a...)
	x := ""
	for i := stack - 1; i >= 0; i-- {
		x += ")"
		a := try(n, k+1, i+1, s+x+"(")
		answer = append(answer, a...)
	}
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
