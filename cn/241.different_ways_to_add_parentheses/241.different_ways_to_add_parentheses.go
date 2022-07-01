//给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。
//
// 生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 10⁴ 。
//
//
//
// 示例 1：
//
//
//输入：expression = "2-1-1"
//输出：[0,2]
//解释：
//((2-1)-1) = 0
//(2-(1-1)) = 2
//
//
// 示例 2：
//
//
//输入：expression = "2*3-4*5"
//输出：[-34,-14,-10,-10,10]
//解释：
//(2*(3-(4*5))) = -34
//((2*3)-(4*5)) = -14
//((2*(3-4))*5) = -10
//(2*((3-4)*5)) = -10
//(((2*3)-4)*5) = 10
//
//
//
//
// 提示：
//
//
// 1 <= expression.length <= 20
// expression 由数字和算符 '+'、'-' 和 '*' 组成。
// 输入表达式中的所有整数值在范围 [0, 99]
//
// 👍 584 👎 0

package cn

import (
	"fmt"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func diffWaysToCompute(expression string) []int {
	//	排列组合
	var nums = make([]int, 0)
	var exp = make([]byte, 0)
	var num = 0
	for i := 0; i < len(expression); i++ {
		x := expression[i]
		switch expression[i] {
		case '+', '-', '*':
			exp = append(exp, x)
			nums = append(nums, num)
			num = 0
		default:
			num = num*10 + int(x-'0')
		}
	}
	nums = append(nums, num)
	var m = make(map[string][]int)
	//var exist = make(map[int]bool)
	cals := Cal(nums, exp, m)
	fmt.Println(m)
	//for i, b := range exist {
	//
	//}
	return cals
}
func Cal(nums []int, exp []byte, m map[string][]int) []int {
	if len(exp) == 0 {
		return nums
	}
	key := toString(nums, exp)
	if v, ok := m[key]; ok {
		return v
	}
	var answer = make([]int, 0)
	// 多少符号，就有多少次运算
	for i := 0; i < len(exp); i++ {
		nn, ne := Cali(nums, exp, i)
		key1 := toString(nn, ne)
		fmt.Println(key, "---", i+1, "--->", key1)
		r := Cal(nn, ne, m)
		for j := 0; j < len(r); j++ {
			answer = append(answer, r[j])
		}
	}
	m[key] = answer
	return m[key]
}

func toString(nums []int, exp []byte) string {
	var str string
	for i := 0; i < len(exp); i++ {
		str += strconv.Itoa(nums[i]) + string([]byte{exp[i]})
	}
	str += strconv.Itoa(nums[len(nums)-1])
	return str
}

func Cali(nums []int, exp []byte, i int) ([]int, []byte) {
	res := getResult(nums[i], nums[i+1], exp[i])
	var nn = make([]int, 0)
	var ne = make([]byte, 0)
	nn = append(nn, nums[:i]...)
	nn = append(nn, res)
	nn = append(nn, nums[i+2:]...)
	ne = append(ne, exp[:i]...)
	ne = append(ne, exp[i+1:]...)
	return nn, ne
}

func getResult(a, b int, exp byte) int {
	switch exp {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
