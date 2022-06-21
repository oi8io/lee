//实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
//
//
//
// BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出
//。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
// boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
// int next()将指针向右移动，然后返回指针处的数字。
//
//
// 注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
//
//
//
// 你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。
//
//
//
// 示例：
//
//
//输入
//["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next",
//"hasNext", "next", "hasNext"]
//[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
//输出
//[null, 3, 7, true, 9, true, 15, true, 20, false]
//
//解释
//BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
//bSTIterator.next();    // 返回 3
//bSTIterator.next();    // 返回 7
//bSTIterator.hasNext(); // 返回 True
//bSTIterator.next();    // 返回 9
//bSTIterator.hasNext(); // 返回 True
//bSTIterator.next();    // 返回 15
//bSTIterator.hasNext(); // 返回 True
//bSTIterator.next();    // 返回 20
//bSTIterator.hasNext(); // 返回 False
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 10⁵] 内
// 0 <= Node.val <= 10⁶
// 最多调用 10⁵ 次 hasNext 和 next 操作
//
//
//
//
// 进阶：
//
//
// 你可以设计一个满足下述条件的解决方案吗？next() 和 hasNext() 操作均摊时间复杂度为 O(1) ，并使用 O(h) 内存。其中 h 是树的高
//度。
//
// 👍 597 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator struct {
	data    []int
	current int
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	arr := make([]int, 0)
	it := BSTIterator{data: arr, current: -1}
	it.inOrder(root)
	return it
}
func (this *BSTIterator) inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	this.inOrder(root.Left)
	this.data = append(this.data, root.Val)
	this.inOrder(root.Right)
}

func (this *BSTIterator) Next() int {
	this.current++
	return this.data[this.current]
}

func (this *BSTIterator) HasNext() bool {
	c := this.current
	c++
	return c < len(this.data)
}

/*type BSTIterator struct {
	Tree    *TreeNode
	curNode *TreeNode
	stack   []*TreeNode
	top     int
}

func Constructor(root *TreeNode) BSTIterator {
	ins := BSTIterator{
		Tree: root,
		top:  -1,
	}
	return ins
}

func (this *BSTIterator) Next() int {
	//fmt.Println(this.top, this.stack)
	var tree *TreeNode
	cur := this.curNode
	if cur == nil {
		tree = this.Tree
	} else if cur.Right != nil { //右子树中最小值
		tree = cur.Right
	} else {
		if this.top < 0 {
			return 0
		}
		tree = this.stack[this.top]
		for tree != nil && cur == tree.Right { // 如果  当前节点 为父节点的右子节点，则找父节点的父亲节点
			cur = tree
			this.top--
		}
		this.stack = this.stack[:this.top]
		this.curNode = tree
		//fmt.Println(cur, "-------->>", tree)
		return tree.Val
	}
	root := tree
	for root != nil && root.Left != nil {
		this.stack = append(this.stack, root)
		root = root.Left
		this.top++
	}
	this.curNode = root
	return this.curNode.Val
}

func (this *BSTIterator) TreeMiniNum(root *TreeNode) *TreeNode {
	for root != nil && root.Left != nil {
		root = root.Left
	}
	return root
}
func (this *BSTIterator) TreeMaxiNum(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func (this *BSTIterator) HasNext() bool {
	var tree *TreeNode
	cur := this.curNode
	if cur == nil {
		tree = this.Tree
	} else if cur.Right != nil { //右子树中最小值
		tree = cur.Right
	} else {
		if this.top < 0 {
			return false
		}

		top := this.top
		tree = this.stack[top]
		for tree != nil && cur == tree.Right { // 如果  当前节点 为父节点的右子节点，则找父节点的父亲节点
			top--
			cur = tree
			tree = this.stack[top]
			if top == 0 {
				return false
			}
		}
		fmt.Println("CURRENT ++++ ", tree, cur, this.top, this.stack)

	}
	return tree != nil
}
*/
/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
//leetcode submit region end(Prohibit modification and deletion)
