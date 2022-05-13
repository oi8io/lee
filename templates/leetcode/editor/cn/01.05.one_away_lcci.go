//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
// 示例 1:
//
// 输入:
//first = "pale"
//second = "ple"
//输出: True
//
//
//
// 示例 2:
//
// 输入:
//first = "pales"
//second = "pal"
//输出: False
//
// 👍 140 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	n1, n2 := len(first), len(second)
	if abs(n1, n2) >= 2 {
		return false
	}
	for i := 0; i < min(n1, n2); i++ {
		if first[i] != second[i] {
			if n1 > n2 {
				return first[i+1:] == second[i:]
			} else if n1 < n2 {
				return first[i:] == second[i+1:]
			} else {
				if i+1 == n1 { //已经遍历完成了
					return true
				}
				return second[i+1:] == first[i+1:]
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
