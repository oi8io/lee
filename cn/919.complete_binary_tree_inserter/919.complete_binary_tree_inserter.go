//完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。
//
// 设计一种算法，将一个新节点插入到一个完整的二叉树中，并在插入后保持其完整。
//
// 实现 CBTInserter 类:
//
//
// CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
// CBTInserter.insert(int v) 向树中插入一个值为 Node.val == val的新节点 TreeNode。使树保持完全二叉树的状态
//，并返回插入节点 TreeNode 的父节点的值；
// CBTInserter.get_root() 将返回树的头节点。
//
//
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入
//["CBTInserter", "insert", "insert", "get_root"]
//[[[1, 2]], [3], [4], []]
//输出
//[null, 1, 2, [1, 2, 3, 4]]
//
//解释
//CBTInserter cBTInserter = new CBTInserter([1, 2]);
//cBTInserter.insert(3);  // 返回 1
//cBTInserter.insert(4);  // 返回 2
//cBTInserter.get_root(); // 返回 [1, 2, 3, 4]
//
//
//
// 提示：
//
//
// 树中节点数量范围为 [1, 1000]
// 0 <= Node.val <= 5000
// root 是完全二叉树
// 0 <= val <= 5000
// 每个测试用例最多调用 insert 和 get_root 操作 10⁴ 次
//
//
// 👍 129 👎 0

package cn

import (
	. "github.com/oi8io/lee/cn/common"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type CBTInserter struct {
	root  *TreeNode
	count int
}

func Constructor(root *TreeNode) CBTInserter {
	count := GetParent(root)
	return CBTInserter{
		root:  root,
		count: count,
	}
}

func GetParent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cnt := 1
	il := GetParent(root.Left)
	ir := GetParent(root.Right)
	return cnt + il + ir
}

func (this *CBTInserter) Insert(val int) int {
	this.count++
	c := this.count
	var stack = make([]int, 0)
	for {
		c = c >> 1
		if c == 1 {
			break
		}
		stack = append(stack, c)
	}
	parent := this.root
	for len(stack) > 0 {
		i := stack[len(stack)-1]
		if i%2 == 1 {
			parent = parent.Right
		} else {
			parent = parent.Left
		}
		stack = stack[:len(stack)-1]
	}
	if this.count%2 == 1 {
		parent.Right = &TreeNode{Val: val}
	} else {
		parent.Left = &TreeNode{Val: val}
	}
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
//leetcode submit region end(Prohibit modification and deletion)
