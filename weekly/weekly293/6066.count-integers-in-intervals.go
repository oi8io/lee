package weekly293

/**
给你区间的 空 集，请你设计并实现满足要求的数据结构：

新增：添加一个区间到这个区间集合中。
统计：计算出现在 至少一个 区间中的整数个数。
实现 CountIntervals 类：

CountIntervals() 使用区间的空集初始化对象
void add(int leftNode, int rightNode) 添加区间 [leftNode, rightNode] 到区间集合之中。
int count() 返回出现在 至少一个 区间中的整数个数。
注意：区间 [leftNode, rightNode] 表示满足 leftNode <= x <= rightNode 的所有整数 x 。



示例 1：

输入
["CountIntervals", "add", "add", "count", "add", "count"]
[[], [2, 3], [7, 10], [], [5, 8], []]
输出
[null, null, null, 6, null, 8]

解释
CountIntervals countIntervals = new CountIntervals(); // 用一个区间空集初始化对象
countIntervals.add(2, 3);  // 将 [2, 3] 添加到区间集合中
countIntervals.add(7, 10); // 将 [7, 10] 添加到区间集合中
countIntervals.count();    // 返回 6
                           // 整数 2 和 3 出现在区间 [2, 3] 中
                           // 整数 7、8、9、10 出现在区间 [7, 10] 中
countIntervals.add(5, 8);  // 将 [5, 8] 添加到区间集合中
countIntervals.count();    // 返回 8
                           // 整数 2 和 3 出现在区间 [2, 3] 中
                           // 整数 5 和 6 出现在区间 [5, 8] 中
                           // 整数 7 和 8 出现在区间 [5, 8] 和区间 [7, 10] 中
                           // 整数 9 和 10 出现在区间 [7, 10] 中


提示：

1 <= leftNode <= rightNode <= 109
最多调用  add 和 count 方法 总计 105 次
调用 count 方法至少一次
*/

type CountIntervals struct {
	root *segmentTree
}
type segmentTree struct {
	leftVal   int
	rightVal  int
	leftNode  *segmentTree
	rightNode *segmentTree
	count     int
	lazy      int
}

func ConstructorCountIntervals() CountIntervals {
	return CountIntervals{}
}

func (this *CountIntervals) Add(left int, right int) {
	if this.root == nil { // 初始化根节点
		this.root = AddNode(nil, left, right)
	}

}

func AddNode(t *segmentTree, left, right int) *segmentTree {
	if t == nil {
		return BuildNewNode(left, right, nil)
	}
	if left <= t.leftVal && right >= t.rightVal { //直接覆盖
		t = BuildNewNode(left, right, t)
		t.leftNode = nil
		t.rightNode = nil
		return t
	}
	if left < t.leftVal {
		AddNode(t.leftNode, left, right)
		t.count += t.leftNode.lazy
		t.lazy += t.leftNode.lazy
		t.leftNode.lazy = 0
	}
	if right > t.rightVal { // 左子树覆盖
		AddNode(t.rightNode, left, right)
		t.count += t.rightNode.lazy
		t.lazy += t.rightNode.lazy
		t.rightNode.lazy = 0
	}
	return t
}

func BuildNewNode(left int, right int, n *segmentTree) *segmentTree {
	if n == nil {
		n = &segmentTree{}
	}
	n.leftVal = left
	n.rightVal = right
	n.count = right - left + 1
	return n
}

func (this *CountIntervals) Count() int {
	if this.root == nil {
		return 0
	}
	return this.root.count
}

/** 解法1. 已废弃
type CountIntervals struct {
	arr   []Item
	count int
}
type Item struct {
	leftNode  int
	rightNode int
}

func Constructor() CountIntervals {
	return CountIntervals{}
}

func (this *CountIntervals) Add(leftNode int, rightNode int) {
	n := len(this.arr)
	add := rightNode - leftNode + 1
	if len(this.arr) == 0 {
		this.arr = append(this.arr, Item{leftNode, rightNode})
		this.count += add
		return
	}
	if leftNode >= this.arr[n-1].rightNode { // 右边新增
		Printf("右不相交 %v->i[%d,%d]  ", this.arr, leftNode, rightNode)
		if leftNode == this.arr[n-1].rightNode {
			this.arr[n-1].rightNode = rightNode
			add--
		} else {
			this.arr = append(this.arr, Item{leftNode, rightNode})
		}
		this.count += add
		return
	}

	if rightNode <= this.arr[0].leftNode { //左边新增
		Printf("左不相交[%d,%d] <- %v  ", leftNode, rightNode, this.arr)
		if rightNode == this.arr[0].leftNode {
			this.arr[0].leftNode = leftNode
			add--
		} else {
			this.arr = append([]Item{{leftNode, rightNode}}, this.arr...)
		}
		this.count += add
		return
	}
	min, max := this.findPos(leftNode, rightNode)
	fmt.Println("数：：", leftNode, rightNode, this.arr, min, max)
	leftVal, end := 0, len(this.arr)-1
	var started, ended bool
	for i := min; i <= max; i++ {
		l, r := this.arr[i].leftNode, this.arr[i].rightNode
		// 1.被当前项包含
		if l <= leftNode && rightNode <= r {
			Printf("直接返回：<%d,%d>被覆盖%v", leftNode, rightNode, this.arr[i])
			return
		}
		if l > rightNode { //右边不相交

		} else {
			if l >= leftNode { // 中间

			} else if l < leftNode { //左边不相交

			}
		}
		// 2.左边不相交
		if r < leftNode {
			started = true
			leftVal = i
			continue
		}

		//3. 右边不相交
		if l > rightNode {
			end = i
			ended = true
			break
		}
		// 4. left和right在中间
		if leftNode < l && r < rightNode { //情况3
			Printf("情况3:被包含，(%d ,%d) 需要删除： %v", leftNode, rightNode, this.arr[i])
			//在新增区间内
			add -= r - l + 1
			continue
		}
		// 5.
		//	部分在新增区域左边，需要将右边的去掉 ,right一定大于ri
		if l <= leftNode { // new【5，9】 old 【3,6】->【3，9】
			started = true
			leftVal = i - 1
			add -= r - leftNode + 1
			Printf("右边超出部分，更新left:(%d<-%d ,%d) %v", l, leftNode, rightNode, this.arr[i])
			leftNode = l
		} else {
			end = i + 1
			ended = true
			add -= rightNode - l + 1
			Printf("左边超出部分,跟新right::(%d ,%d->%d) %v", leftNode, rightNode, r, this.arr[i])
			rightNode = r
		}
		if leftNode > r {
			Printf("在左边:需要在右侧插入：(%d ,%d->%d) %v", leftNode, rightNode, r, this.arr[i])
		}
	}

	//Printf(leftNode, rightNode, this.arr, this.count, "总计添加：", add, "=", this.count+add)
	//Printf("开始结束打印：   arr:%v leftVal:end:【%d,%d】 [leftVal:end]%v (leftNode,rightNode)(%d,%d)", this.arr, leftVal, end, this.arr[leftVal:end], leftNode, rightNode)
	//this.updateData(leftVal, end, Item{leftNode, rightNode})
	this.count += add
	item := Item{leftNode, rightNode}
	var leftArr []Item
	rightArr := []Item{item}

	if started {
		leftArr = this.arr[:leftVal+1] //右不包含
	}
	if ended {
		rightArr = append(rightArr, this.arr[end:]...)
	}
	this.arr = append(leftArr, rightArr...)
	//Printf(leftVal, end, started, "最终拼接：：：：：：：", this.arr, leftArr, leftNode, rightNode, rightArr)

	//Printf("最终结果::::", this.count, leftNode, rightNode, this.arr)
}

func (this *CountIntervals) updateData(leftVal, end int, item Item) {
	leftArr := this.arr[:leftVal+1] //右不包含
	rightArr := append([]Item{item}, this.arr[end:]...)
	this.arr = append(leftArr, rightArr...)
}

func (this *CountIntervals) findPos(leftNode, rightNode int) (int, int) {
	min, max := 0, len(this.arr)-1
	for {
		var update bool
		if max-min <= rightNode-leftNode {
			break
		}
		cur := (min + max) / 2
		if cur == min || max == cur {
			break
		}
		if this.arr[cur].leftNode >= rightNode { // 右边
			update = true
			max = cur
		}
		if this.arr[cur].rightNode <= leftNode { // 左边
			update = true
			min = cur
		}
		if !update {
			break
		}
	}
	return min, max
}

func (this *CountIntervals) Count() int {
	return this.count
}

func Printf(format string, a ...interface{}) {
	var p = false
	if p {
		fmt.Printf(format, a...)
		fmt.Println()
	}
}
*/
/**
 * Your CountIntervals object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(leftNode,rightNode);
 * param_2 := obj.Count();
 */
