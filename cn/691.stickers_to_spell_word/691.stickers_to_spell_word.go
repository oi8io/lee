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

import (
	"fmt"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minStickers(stickers []string, target string) int {
	//每次尝试一种纸，直到完成，另外就是target中有字符不在stickers中
	//

	//noChangeMap := make(map[string]bool)
	//cntMap := make(map[string]int)
	//min := tryMinStickers(stickers, target, noChangeMap, cntMap)
	min := minStickers3(stickers, target)
	if min == math.MaxInt {
		return -1
	}
	return min
}

// 动态规划
func minStickers3(stickers []string, target string) int {
	//1. 先整理成二维数组
	var stickersArr = make([][]int, len(stickers), len(stickers))
	for i := 0; i < len(stickers); i++ {
		stickersArr[i] = stringToCharArr(stickers[i])
	}
	var exists = make(map[string]int)
	//先排序
	target = charArrToString(stringToCharArr(target))
	stickers3 := tryMinStickers3(stickersArr, target, exists)
	return stickers3
}

func stringToCharArr(s string) []int {
	var a = make([]int, 26, 26)
	for i := 0; i < len(s); i++ {
		a[s[i]-'a']++
	}
	return a
}
func charArrToString(a []int) string {
	var sa []byte
	for i := 0; i < len(a); i++ {
		for j := 0; j < a[i]; j++ {
			sa = append(sa, byte(i+'a'))
		}
	}
	return string(sa)
}

// 通过同位比较，大就直接减到0，小就减去对应位置的数字
func minusArr(s []int, p []int) string {
	var ret = make([]int, 26, 26)
	copy(ret, s)
	for i := 0; i < len(p); i++ {
		if s[i] > 0 {
			if s[i] > p[i] {
				ret[i] = s[i] - p[i]
			} else {
				ret[i] = 0
			}
		}
	}
	return charArrToString(ret)
}

func tryMinStickers3(stickers [][]int, target string, exist map[string]int) int {
	if len(target) == 0 {
		return 0
	}
	if v, ok := exist[target]; ok {
		return v
	}
	var targetArr = stringToCharArr(target)
	min := math.MaxInt
	for i := 0; i < len(stickers); i++ {
		// 首字母匹配的里面去挑选
		if stickers[i][target[0]-'a'] > 0 {
			t := minusArr(targetArr, stickers[i])
			ret := tryMinStickers3(stickers, t, exist)
			if ret != math.MaxInt && ret+1 < min {
				min = ret + 1
			}
		}
	}
	exist[target] = min
	return min
}

// 回溯，暴力递归
func minStickers2(stickers []string, target string) int {
	//每次尝试一种纸，直到完成，另外就是target中有字符不在stickers中
	//

	noChangeMap := make(map[string]bool)
	cntMap := make(map[string]int)

	min := tryMinStickers(stickers, target, noChangeMap, cntMap)
	if min == math.MaxInt {
		return -1
	}
	return min
}

// 回溯，暴力递归
func tryMinStickers(stickers []string, target string, noChangeMap map[string]bool, cntMap map[string]int) int {
	if len(target) == 0 {
		return 0
	}
	fmt.Println(target)
	if v, ok := cntMap[target]; ok {
		return v
	}

	//fmt.Println(targetMap)
	//每次尝试一种纸，直到完成，另外就是target中有字符不在stickers中
	min := math.MaxInt
	for i := 0; i < len(stickers); i++ {
		key := target + "++" + stickers[i]
		if _, ok := noChangeMap[key]; ok {
			continue
		}
		m, changed := minus(stickers[i], target)
		if changed {
			i2 := tryMinStickers(stickers, m, noChangeMap, cntMap)
			if i2 != math.MaxInt && i2+1 < min {
				min = i2 + 1
			}
		} else {
			noChangeMap[key] = true
		}
	}
	cntMap[target] = min
	return min
}

func minus(sticker string, target string) (string, bool) {
	var targetMap = make(map[byte]int)
	for i, _ := range target {
		c := target[i]
		targetMap[c]++
	}
	var changed bool
	for n := 0; n < len(sticker); n++ {
		c := sticker[n]
		if _, ok := targetMap[c]; ok {
			changed = true
			targetMap[c]--
			if targetMap[c] <= 0 {
				delete(targetMap, c)
			}
		}
	}
	if !changed {
		return target, changed
	}
	var ret []byte
	for i := 0; i < len(target); i++ {
		c := target[i]
		if i2, ok := targetMap[c]; ok && i2 > 0 {
			ret = append(ret, c)
			targetMap[c]--
		}
	}
	return string(ret), changed

}

func minStickers1(stickers []string, target string) int {

	var targetMap = make(map[byte]int)
	for i, _ := range target {
		c := target[i]
		targetMap[c]++
	}
	var countMap = make(map[byte][]int)
	var wordCount = make([]map[byte]int, len(stickers))
	for i, sticker := range stickers {
		wordCount[i] = make(map[byte]int)
		for n, _ := range sticker {
			c := sticker[n]
			/*			if _, ok := targetMap[c]; !ok { //不需要的字符忽略掉
						continue
					}*/
			if _, ok := countMap[c]; !ok {
				countMap[c] = make([]int, len(stickers))
			}
			countMap[c][i]++
		}
	}
	var reslovers = make(map[byte]map[int]int)
	for c, cnt := range targetMap {
		if _, ok := countMap[c]; !ok {
			return -1
		}
		reslovers[c] = matchSticker(cnt, countMap[c])
	}
	var min = len(targetMap)
	for c, _ := range targetMap {
		if c != 97 {
			continue
		}
		for i, _ := range reslovers[c] {
			nMap := copyMap(targetMap)
			n := trySticker(c, i, nMap, 0, countMap, wordCount, reslovers)
			//fmt.Printf(c, i, n)
			if n < min {
				min = n
			}
		}
	}
	return min
}

func trySticker(c byte, i int, cMap map[byte]int, words int, countMap map[byte][]int, wordCount []map[byte]int, reslovers map[byte]map[int]int) int {
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

func copyMap(cMap map[byte]int) map[byte]int {
	var newMap = make(map[byte]int)
	for ch, num := range cMap {
		newMap[ch] = num
	}
	return newMap
}
func copyMap1(cMap map[byte]int) map[byte]int {
	var newMap = make(map[byte]int)
	for ch, num := range cMap {
		newMap[ch] = num
	}
	return newMap
}

func DealMap(cMap, wordCount map[byte]int, n int) map[byte]int {
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
