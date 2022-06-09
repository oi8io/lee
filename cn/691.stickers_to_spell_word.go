//我们有 n 种不同的贴纸。每个贴纸上都有一个小写的英文单词。
//
// 您想要拼写出给定的字符串 target ，方法是从收集的贴纸中切割单个字母并重新排列它们。如果你愿意，你可以多次使用每个贴纸，每个贴纸的数量是无限的。
//
// 返回你需要拼出 target 的最小贴纸数量。如果任务不可能，则返回 -1 。
//
// 注意：在所有的测试用例中，所有的单词都是从 1000 个最常见的美国英语单词中随机选择的，并且 target 被选择为两个随机单词的连接。
//
//
//
// 示例 1：
//
//
//输入： stickers = ["with","example","science"], target = "thehat"
//输出：3
//解释：
//我们可以使用 2 个 "with" 贴纸，和 1 个 "example" 贴纸。
//把贴纸上的字母剪下来并重新排列后，就可以形成目标 “thehat“ 了。
//此外，这是形成目标字符串所需的最小贴纸数量。
//
//
// 示例 2:
//
//
//输入：stickers = ["notice","possible"], target = "basicbasic"
//输出：-1
//解释：我们不能通过剪切给定贴纸的字母来形成目标“basicbasic”。
//
//
//
// 提示:
//
//
// n == stickers.length
// 1 <= n <= 50
// 1 <= stickers[i].length <= 10
// 1 <= target <= 15
// stickers[i] 和 target 由小写英文单词组成
//
// 👍 197 👎 0

package cn

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func minStickers(stickers []string, target string) int {

	var cMap = make(map[int32]int)
	for _, c := range target {
		cMap[c]++
	}
	var countMap = make(map[int32][]int)
	var wordCount = make([]map[int32]int, len(stickers))
	for i, sticker := range stickers {
		wordCount[i] = make(map[int32]int)
		for _, c := range sticker {
			if _, ok := cMap[c]; !ok { //不需要的字符忽略掉
				continue
			}
			if _, ok := countMap[c]; !ok {
				countMap[c] = make([]int, len(stickers))
			}
			countMap[c][i]++
		}
	}
	var reslovers = make(map[int32]map[int]int)
	for c, cnt := range cMap {
		if _, ok := countMap[c]; !ok {
			return -1
		}
		reslovers[c] = matchSticker(cnt, countMap[c])
	}
	var min = len(cMap)
	for c, _ := range cMap {
		if c != 97 {
			continue
		}
		for i, _ := range reslovers[c] {
			nMap := copyMap(cMap)
			n := trySticker(c, i, nMap, 0, countMap, wordCount, reslovers)
			//fmt.Printf(c, i, n)
			if n < min {
				min = n
			}
		}
	}
	return min
}

func trySticker(c int32, i int, cMap map[int32]int, words int, countMap map[int32][]int, wordCount []map[int32]int, reslovers map[int32]map[int]int) int {
	if len(cMap) == 0 {
		return words
	}
	var min = len(reslovers)
	n := reslovers[c][i]
	if n == 0 {
		return words
	}
	delete(cMap, c)
	cMap = DealMap(cMap, wordCount[i], n)
	fmt.Println(cMap, words, n, min)
	words = words + n
	if words > 0 && words < min {
		min = words
	}
	if len(cMap) == 0 {
		return min
	}
	for c2, _ := range cMap {
		for i2, _ := range reslovers[c] {
			words = trySticker(c2, i2, cMap, words, countMap, wordCount, reslovers)
			if words > 0 && words < min {
				min = words
			}
		}
	}

	fmt.Println("ENDDDDD", cMap, words, n, min)
	return min
}

func copyMap(cMap map[int32]int) map[int32]int {
	var newMap = make(map[int32]int)
	for ch, num := range cMap {
		newMap[ch] = num
	}
	return newMap
}

func DealMap(cMap, wordCount map[int32]int, n int) map[int32]int {
	for ch, num := range wordCount {
		if cMap[ch] > num*n {
			cMap[ch] = cMap[ch] - num*n
		} else {
			delete(cMap, ch)
		}
	}
	return cMap
}

// 每个字母有多少种方法
func matchSticker(cnt int, counts []int) map[int]int {
	var solvers = make(map[int]int)
	for i := 0; i < len(counts); i++ {
		if counts[i] > 0 {
			solvers[i] = cnt / counts[i]
		}
	}
	return solvers
}

//leetcode submit region end(Prohibit modification and deletion)
