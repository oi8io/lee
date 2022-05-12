/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 *
 * https://leetcode-cn.com/problems/remove-duplicate-letters/description/
 *
 * algorithms
 * Medium (47.81%)
 * Likes:    704
 * Dislikes: 0
 * Total Accepted:    83.9K
 * Total Submissions: 175.6K
 * Testcase Example:  '"bcabc"'
 *
 * 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "bcabc"
 * 输出："abc"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbacdcbc"
 * 输出："acdb"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 由小写英文字母组成
 *
 *
 *
 *
 * 注意：该题与 1081
 * https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters
 * 相同
 *
 */

package problems

import "fmt"

// @lc code=start
/**
cbacdcbc
cba c
b<c -> bac
bacd c
bacd
bacd b
b>a acdb
acdb c



bcabc

b
bc
bca
bca b -> cab
cab c -> abc



*/
func removeDuplicateLetters1(s string) string {
	fmt.Println(s)
	var existsChar = make(map[byte]int)
	var ret []byte
	for j := len(s) - 1; j >= 0; j-- {
		c := s[j]
		if _, ok := existsChar[c]; !ok {
			ret = append([]byte{c}, ret...)
			existsChar[c] = j
		}
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		fmt.Printf("%5s ", ret)
		if _, ok := existsChar[c]; !ok {
			ret = append([]byte{c}, ret...)
			existsChar[c] = 10
			continue
		}

		i := 0
		for ; i < len(ret); i++ {
			if c == ret[i] {
				break
			}
		}
		//第一个字符，直接跳过
		if i == 0 {
			continue
		}
		// 如果后面的字母要大，则删除当前字符，再追加
		if ret[0] > ret[i] {
			copy(ret[i:], ret[i+1:])
			ret[len(ret)-1] = c
		}
		fmt.Printf("[%c] %s  \n", c, ret)
	}

	return string(ret)
}

// @lc code=end

func removeDuplicateLetters(s string) string {
	left := [26]int{} // 统计出现的次数
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		num := ch - 'a'

		if inStack[num] {
			left[num]--
			continue
		}
		// 栈不为空 ，小于栈顶元素
		for len(stack) > 0 && ch < stack[len(stack)-1] {
			last := stack[len(stack)-1] - 'a' // 判断后面还有没有该字符
			if left[last] == 0 {              // 如果栈顶元素没有剩余了
				break
			}
			//如果后面还有，移除栈顶元素
			stack = stack[:len(stack)-1]
			inStack[last] = false //等着后续再加进来
		}
		stack = append(stack, ch)
		inStack[num] = true
		left[num]--
	}
	return string(stack)
}
