//给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。
//
// 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序。
//
//
//
// 示例 1：
//
//
//输入: words = ["i", "love", "leetcode", "i", "love", "coding"], k = 2
//输出: ["i", "love"]
//解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。
//
//
// 示例 2：
//
//
//输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
//k = 4
//输出: ["the", "is", "sunny", "day"]
//解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//    出现次数依次为 4, 3, 2 和 1 次。
//
//
//
//
// 注意：
//
//
// 1 <= words.length <= 500
// 1 <= words[i] <= 10
// words[i] 由小写英文字母组成。
// k 的取值范围是 [1, 不同 words[i] 的数量]
//
//
//
//
// 进阶：尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
// 👍 470 👎 0

package cn

import (
	"container/heap"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)

type kv struct {
	k string
	v int
}

type sh []kv

func (h sh) Len() int {
	return len(h)
}

func (h sh) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.v < b.v || (a.v == b.v && a.k > b.k)
}

func (h sh) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *sh) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

func (h *sh) Pop() (x interface{}) {
	a := *h
	x = a[h.Len()-1]
	*h = a[:h.Len()-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	return topKFrequent2(words, k)
	return topKFrequent1(words, k)
}
func topKFrequent2(words []string, k int) []string {
	var m = make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	var h = &sh{}
	for s, i := range m {
		heap.Push(h, kv{s, i})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(kv).k
	}
	return ans
}

func topKFrequent1(words []string, k int) []string {
	var m = make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	type kv struct {
		v int
		k string
	}
	var x = make([]kv, 0)
	for s, i := range m {
		x = append(x, kv{i, s})
	}
	sort.Slice(x, func(i, j int) bool {
		a, b := x[i], x[j]
		return a.v > b.v || (a.v == b.v && a.k < b.k)
	})
	var ans = make([]string, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, x[i].k)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
