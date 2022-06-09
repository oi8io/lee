//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
//
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
//
//
//
//
// 示例 1:
//
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// -2³¹ <= val <= 2³¹ - 1
// pop、top 和 getMin 操作总是在 非空栈 上调用
// push, pop, top, and getMin最多被调用 3 * 10⁴ 次
//
// 👍 1301 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	data  []int
	min   int
	empty bool
	index int
}

func ConstructorMinStack() MinStack {
	return MinStack{data: make([]int, 1), index: -1, empty: true}
}

func (this *MinStack) Push(val int) {
	if this.empty {
		this.min = val
		this.empty = false
	} else {
		if val < this.min {
			this.min = val
		}
	}
	this.index++
	if len(this.data)-1 < this.index {
		this.data = append(this.data, make([]int, this.index)...)
	}
	this.data[this.index] = val
}

func (this *MinStack) Pop() {
	if this.Top() == this.min {
		min := this.data[0]
		for i := 0; i < this.index; i++ {
			if this.data[i] < min {
				min = this.data[i]
			}
		}
		this.min = min
	}
	this.index--
	if this.index == -1 {
		this.empty = true
	}
}

func (this *MinStack) Top() int {
	return this.data[this.index]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
