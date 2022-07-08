package heap

type Heap interface {
	isEmpty() bool  //是否为空
	heapify(i int)  //从元素开始，与孩子比较，如果比孩子大（小），则取大（小）进行交换，然后再进入下一层进行比较
	push(n int)     // 尾部插入元素，再将给元素与父亲节点比较，如果大（小）则交换位置继续比较
	pop() int       // 弹出根，将最后一个元素放入该位置，然后执行heapify
	getData() []int // 返回所有数据
	top() int       // 返回根
	size() int      //返回大小
}

// NewHeap 新建堆。IsBigTop 为true则大顶堆
func NewHeap(arr []int, IsBigTop bool) (h Heap) {
	if IsBigTop {
		h = &heapBigTop{0, make([]int, 1)}
	} else {
		h = &heapSmallTop{0, make([]int, 1)}
	}
	for i := 0; i < len(arr); i++ {
		h.push(arr[i])
	}
	return h
}
