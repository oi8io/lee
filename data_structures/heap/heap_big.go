package heap

import (
	"math"
)

type heapBigTop struct {
	n    int
	data []int
}

func (h *heapBigTop) getData() []int {
	return h.data[1:]
}

func (h *heapBigTop) isEmpty() bool {
	return h.n == 0
}

func (h *heapBigTop) heapify(i int) {
	for i<<1 <= h.n {
		left, right, next := h.data[i<<1], math.MinInt, i<<1
		if i<<1+1 <= h.n { //如果右子节点不存在，则比对子左节点
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
	return h.n
}

func (h *heapBigTop) top() int {
	return h.data[1]
}

// 插入元素 尾部插入，然后再比较父节点
func (h *heapBigTop) push(n int) {
	h.data = append(h.data, n)
	h.n++
	i := h.n
	// 如果当前元素比父亲节点还大，则交换，否则就退出
	for i>>1 >= 1 && h.data[i] > h.data[i>>1] {
		h.swap(i>>1, i)
		i = i >> 1
	}
}

// 弹出顶部元素，将最后元素替换头再平衡
func (h *heapBigTop) pop() int { // 最大元素
	x := h.data[1]
	h.data[1] = h.data[h.n]
	h.data = h.data[:h.n]
	h.n--
	h.heapify(1)
	return x
}
