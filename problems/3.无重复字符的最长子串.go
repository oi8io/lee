package problems

import "fmt"

/**
解题思路

1. 对字符串进行遍历，将当前字符加入到map中
1. 如果map中已经存在当前字符，而且重复位置在起始位置之后，则字符串被分割成两个部分
1. 前面部分为从起始位置到之前重复字符的位置，后面部分由重复字符的下一个位置开始计算
1. 变更起始位置到字符字符的下一个字符，重新开始计数
1. 其他情况计数加1，并将字符更新到map中
*/

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	var x = make(map[uint8]int)
	var start, max, current = 0, 0, 0

	for i := 0; i < len(s); i++ {
		if n, ok := x[s[i]]; ok && n >= start {
			// n 的位置曾经出现过该字符
			fmt.Printf("before:%d,now:%d->repeat char:%c,current:%d,max:%d\n", n, i, s[i], current, max)
			start = n
			x[s[i]] = i
			current = i - n //当前总计1个字符
		} else {
			current++
			x[s[i]] = i
			fmt.Printf("append %c,po:%d\n", s[i], current)
		}

		if current >= max {
			max = current
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)

func lengthOfLongestSubstringBefore(s string) int {
	longest := make(map[int]int)
	longest[0] = 1
	var sub string
	for i := 0; i < len(s); i++ {
		if i == 0 {
			sub = string(s[i])
			longest[i] = 1
		} else {
			sub = existsChar(sub, s[i])
			longest[i] = Max(longest[i-1], len(sub))
		}
	}
	return longest[len(s)-1]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func existsChar(s string, c uint8) string {
	for j := 0; j < len(s); j++ {
		if s[j] == c {
			return s[j+1:] + string(c)
		}
	}
	return s + string(c)
}
