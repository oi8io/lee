//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。 
//
// 请实现 KthLargest 类： 
//
// 
// KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。 
// int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["KthLargest", "add", "add", "add", "add", "add"]
//[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
//输出：
//[null, 4, 5, 5, 8, 8]
//
//解释：
//KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
//kthLargest.add(3);   // return 4
//kthLargest.add(5);   // return 5
//kthLargest.add(10);  // return 5
//kthLargest.add(9);   // return 8
//kthLargest.add(4);   // return 8
// 
//
// 
//
// 提示： 
//
// 
// 1 <= k <= 10⁴ 
// 0 <= nums.length <= 10⁴ 
// -10⁴ <= nums[i] <= 10⁴ 
// -10⁴ <= val <= 10⁴ 
// 最多调用 add 方法 10⁴ 次 
// 题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素 
// 
//
// 
//
// 注意：本题与主站 703 题相同： https://leetcode-cn.com/problems/kth-largest-element-in-a-
//stream/ 
// 👍 31 👎 0


package cn

import . "github.com/oi8io/lee/cn/common"

//leetcode submit region begin(Prohibit modification and deletion)
// KthLargest 小顶堆
type KthLargest struct {
	root *TreeNode
	max  int
	min  int
	k    int
	len  int
}

func Constructor(k int, nums []int) KthLargest {
	c := KthLargest{k: k, len: 0}
	for i := 0; i < len(nums); i++ {
		c.Add(nums[i])
	}
	return c
}

func (this *KthLargest) Add(val int) int {
	if this.root == nil {
		this.root = &TreeNode{Val: val}
		this.len++
		this.min = val
		return val
	}
	if this.isFull() {
		if val <= this.min { // 比min小，舍弃，
			return this.min
		}
		/*		if found := this.search(this.root, val); found != nil {
				return this.min
			}*/
		this.root = this.deleteMin(this.root)
		this.root = this.insert(this.root, val)
		return this.getMin()
	} else {
		/*		if found := this.search(this.root, val); found != nil {
				return this.getMin()
			}*/
		this.root = this.insert(this.root, val)
		this.len++
		return this.getMin()
	}
}

func (this *KthLargest) isFull() bool {
	return this.k == this.len
}

func (this *KthLargest) deleteMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left = this.deleteMin(root.Left)
		return root
	} else {

	}
	return root.Right
}

func (this *KthLargest) search(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return this.search(root.Left, val)
	}
	if root.Val < val {
		return this.search(root.Right, val)
	}
	return nil
}

func (this *KthLargest) insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val >= val {
		root.Left = this.insert(root.Left, val)
		return root
	}
	if root.Val < val {
		root.Right = this.insert(root.Right, val)
		return root
	}
	return root
}

func (this *KthLargest) getMin() int {
	tree := this.root
	for tree != nil && tree.Left != nil {
		tree = tree.Left
	}
	this.min = tree.Val
	return tree.Val
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)
