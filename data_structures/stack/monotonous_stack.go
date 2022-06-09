package stack

//单调栈 在复杂度O（n）时间复杂度计算数组元素在某个区间最大/最小

// [1,2,3,9,2,1,6]
// 0->1 [-1,-1]
// 1->2 [0,5]
// 2->3 [1,5]
// 3->9 [2,5]
// 4->2 [0,5]
// 5->1 [-1,-1]
// 6->6 [5,-1]

type MonoStack interface {
	peek() []int
	pop() []int
	push(int)
	isEmpty() bool
}

type monoStack struct {
	data [][]int
	top  int
}

func newMonoStack() *monoStack {
	return &monoStack{data: make([][]int, 0)}
}

func (m *monoStack) peek() []int {
	if len(m.data) == 0 {
		return nil
	}
	return m.data[m.top]
}

func (m *monoStack) pop() []int {
	data := m.data[m.top]
	m.data = m.data[:m.top]
	m.top--
	return data
}

func (m *monoStack) add(i int) {
	m.data[m.top] = append(m.data[m.top], i)
}

func (m *monoStack) push(i int) {
	m.data = append(m.data, []int{i})
	m.top++
}
func (m *monoStack) isEmpty() bool {
	return len(m.data) > m.top
}

func MakeMonotonousStack(arr []int) [][]int {
	var res, stack = make([][]int, 0), newMonoStack()
	for i := 0; i < len(arr); i++ {
		// i -> arr[i]
		for !stack.isEmpty() && arr[stack.peek()[0]] > arr[i] {
			popIs := stack.pop()
			leftLessIndex := -1
			if !stack.isEmpty() { // 上一个数字的最后一个位置为前面小于该数字的数
				leftLessIndex = stack.peek()[len(stack.peek())-1]
			}
			for i := 0; i < len(popIs); i++ { //出列的需要计算位置，所有的右边
				res[i][0] = leftLessIndex
				res[i][1] = i
			}
		}
		// 1) 弹出完之后，顶部有东西，和当前值一样
		if !stack.isEmpty() && arr[stack.peek()[0]] == arr[i] {
			stack.add(i)
		} else { // 2) 栈为空  or  栈顶值 < 当前值
			stack.push(i)
		}
	}
	// 单独清算栈里的东西！
	for !stack.isEmpty() {
		popIs := stack.pop()
		leftLessIndex := -1
		if !stack.isEmpty() { // 上一个数字的最后一个位置为前面小于该数字的数
			leftLessIndex = stack.peek()[len(stack.peek())-1]
		}
		for i := 0; i < len(popIs); i++ { // 出列的需要计算位置，所有的右边
			res[i][0] = leftLessIndex
			res[i][1] = -1
		}
	}
	return res

}
