//设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
//
// 示例：
//
// 输入： arr = [1,3,5,7,2,4,6,8], k = 4
//输出： [1,2,3,4]
//
//
// 提示：
//
//
// 0 <= len(arr) <= 100000
// 0 <= k <= min(100000, len(arr))
//
// 👍 189 👎 0

package cn

import (
	"container/heap"
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func smallestK(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 || k == 0 {
		return nil
	}

	return smallestK1(arr, k)
	return smallestK3(arr, k)
	return smallestK2(arr, k)
}

type mh struct {
	sort.IntSlice
}

func (m mh) Push(x interface{}) {}

func (m mh) Pop() (_ interface{}) { return }

func smallestK4(arr []int, k int) []int {
	var x sort.IntSlice = arr
	l, r := 0, x.Len()-1
	m := (l + r) / 2
	if x.Less(m, l) { //左边比中间大
		x.Swap(m, l)
	}
	if x.Less(r, l) { //右边比左边大
		x.Swap(r, l)
		if x.Less(r, m) {
			x.Swap(r, m)
		}
	}
	sort.Ints()

}
func smallestK3(arr []int, k int) []int {
	h := mh{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if v < h.IntSlice[0] {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}
func smallestK2(arr []int, k int) []int {
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}
func smallestK1(arr []int, k int) []int {
	heap := NewHeap(arr, k, true)
	return heap.getData()
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

type Heap interface {
	isEmpty() bool  //是否为空
	heapify(i int)  //从元素开始，与孩子比较，如果比孩子大（小），则取大（小）进行交换，然后再进入下一层进行比较
	push(n int)     // 尾部插入元素，再将给元素与父亲节点比较，如果大（小）则交换位置继续比较
	pop() int       // 弹出根，将最后一个元素放入该位置，然后执行heapify
	getData() []int // 返回所有数据
	top() int       // 返回根
	size() int      //返回大小
	cap() int       //返回大小
}

// NewHeap 新建堆。IsBigTop 为true则大顶堆
func NewHeap(arr []int, cap int, IsBigTop bool) (h Heap) {
	if IsBigTop {
		h = &heapBigTop{0, cap, make([]int, 1)}
	} else {
		h = &heapBigTop{0, cap, make([]int, 1)}
	}
	for i := 0; i < len(arr); i++ {
		h.push(arr[i])
	}
	return h
}

type heapBigTop struct {
	len  int
	capt int
	data []int
}

func (h *heapBigTop) cap() int {
	return h.capt
}

func (h *heapBigTop) getData() []int {
	return h.data[1:]
}

func (h *heapBigTop) isEmpty() bool {
	return h.len == 0
}

func (h *heapBigTop) heapify(i int) {
	for i<<1 <= h.len {
		left, right, next := h.data[i<<1], math.MinInt, i<<1
		if i<<1+1 <= h.len { //如果右子节点不存在，则比对子左节点
			right = h.data[i<<1+1]
		}
		if left < right {
			next = i<<1 + 1
		}
		if h.data[i] > h.data[next] {
			break
		}
		h.swap(next, i)
		i = next
	}
}

func (h *heapBigTop) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *heapBigTop) size() int {
	return h.len
}

func (h *heapBigTop) top() int {
	return h.data[1]
}

func (h *heapBigTop) isFull() bool {
	return h.len == h.capt
}

// 插入元素 尾部插入，然后再比较父节点
func (h *heapBigTop) push(n int) {
	if h.isFull() {
		if n < h.top() {
			h.data[0] = n
			h.heapify(0)
		}
		return
	}
	h.data = append(h.data, n)
	h.len++
	i := h.len
	// 如果当前元素比父亲节点还大，则交换，否则就退出
	for i>>1 >= 1 && h.data[i] > h.data[i>>1] {
		h.swap(i>>1, i)
		i = i >> 1
	}
}

// 弹出顶部元素，将最后元素替换头再平衡
func (h *heapBigTop) pop() int { // 最大元素
	x := h.data[1]
	h.data[1] = h.data[h.len]
	h.data = h.data[:h.len]
	h.len--
	h.heapify(1)
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
