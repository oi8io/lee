//// 基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
////
//// 假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
////
////
//// 例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
////
////
//// 另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。
////
//// 给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法
//完成此基因变化，返回 -1 。
////
//// 注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。
////
////
////
//// 示例 1：
////
////
////输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
////输出：1
////
////
//// 示例 2：
////
////
////输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA",
//"AAACGGTA"]
////输出：2
////
////
//// 示例 3：
////
////
////输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC",
//"AACCCCCC"]
////输出：3
////
////
////
////
//// 提示：
////
////
//// start.length == 8
//// end.length == 8
//// 0 <= bank.length <= 10
//// bank[i].length == 8
//// start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成
////
//// 👍 161 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func minMutation(start string, end string, bank []string) int {
	//回溯算法
	//1. 找到start相差1个字符的字符串，选择一个进行下一步，并且记录访问过的字符串，如果字符串与目标相等，则返回，如果没有能访问的数据库之后，返回失败
	hasVisited := 0
	mutation, _ := mutation(start, end, 0, bank, hasVisited)
	return mutation
}

func mutation(start, end string, n int, bank []string, hasVisited int) (int, bool) {
	//fmt.Println(start)
	var result []int
	var hasNext, success bool
	for i := 0; i < len(bank); i++ {
		cur := bank[i]
		if 1<<i&hasVisited != 0 {
			continue
		}
		if cur == start {
			hasVisited = hasVisited | 1<<i
			continue
		}
		diff := 0
		for j := 0; j < 8; j++ {
			if start[j] != cur[j] {
				diff++
			}
		}
		if diff == 1 {
			hasNext = true
			//fmt.Println(start, cur)
			newVisited := hasVisited | 1<<i
			if cur == end {
				return n + 1, true
			}
			last, b := mutation(cur, end, n+1, bank, newVisited)
			if b {
				success = true
				result = append(result, last)
			}
		}
	}

	if !hasNext || !success {
		return -1, false
	}
	min := len(bank)
	for _, i2 := range result {
		if i2 < min {
			min = i2
		}
	}
	return min, true
}

//leetcode submit region end(Prohibit modification and deletion)
