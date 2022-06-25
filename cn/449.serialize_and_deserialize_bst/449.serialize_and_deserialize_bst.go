//序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
//
// 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序
//列化为最初的二叉搜索树。
//
// 编码的字符串应尽可能紧凑。
//
//
//
// 示例 1：
//
//
//输入：root = [2,1,3]
//输出：[2,1,3]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数范围是 [0, 10⁴]
// 0 <= Node.val <= 10⁴
// 题目数据 保证 输入的树是一棵二叉搜索树。
//
// 👍 363 👎 0

package cn

import (
	"strconv"
	"strings"
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type OrderType string

var (
	preOrder   OrderType = "preOrder"
	inOrder    OrderType = "inOrder"
	postOrder  OrderType = "postOrder"
	levelOrder OrderType = "levelOrder"
)

type Codec struct {
	data  []string
	order OrderType
}

func Constructor() Codec {
	return Codec{data: make([]string, 0), order: levelOrder}
}

func (this *Codec) front() *string {
	if len(this.data) == 0 {
		return nil
	}
	v := this.data[0]
	this.data = this.data[1:]
	return &v
}

func (this *Codec) push(str string) {
	this.data = append(this.data, str)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 线序遍历
	switch this.order {
	case preOrder:
		this.serializePre(root)
	case inOrder:
		this.serializeIn(root)
	case postOrder:
		this.serializePost(root)
	case levelOrder:
		this.serializeLevel(root)
	}
	return strings.Join(this.data, ",")
}

func (this *Codec) serializePre(root *TreeNode) {
	// 前序遍历
	if root == nil {
		this.data = append(this.data, "null")
		return
	}
	this.data = append(this.data, strconv.Itoa(root.Val))
	this.serializePre(root.Left)
	this.serializePre(root.Right)
	return
}

// Serializes a tree to a single string.
func (this *Codec) serializeIn(root *TreeNode) {

}
func (this *Codec) serializePost(root *TreeNode) {

}

/**

 */
func (this *Codec) serializeLevel(head *TreeNode) {
	if head == nil {
		this.data = append(this.data, "null")
		return
	}
	var queue = make([]*TreeNode, 0)
	queue = append(queue, head)
	for len(queue) > 0 {
		head = queue[0]
		queue = queue[1:]
		if head == nil {
			this.data = append(this.data, "null")
		} else {
			this.data = append(this.data, strconv.Itoa(head.Val))
			queue = append(queue, head.Left)
			queue = append(queue, head.Right)
		}
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" || data == "null" {
		return nil
	}
	this.data = strings.Split(data, ",")
	switch this.order {
	case preOrder:
		return this.deserializePre()
	case inOrder:
		return this.deserializeIn()
	case postOrder:
		return this.deserializePost()
	case levelOrder:
		return this.deserializeLevel()
	}
	return nil
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializePre() *TreeNode {
	if len(this.data) == 0 {
		return nil
	}
	v := this.data[0]
	this.data = this.data[1:] //pop
	if v == "null" {
		return nil
	}
	val, _ := strconv.Atoi(v)
	root := &TreeNode{Val: val}
	root.Left = this.deserializePre()
	root.Right = this.deserializePre()
	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeIn() *TreeNode {
	return nil
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializePost() *TreeNode {
	return nil
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeLevel() *TreeNode {
	if len(this.data) == 0 {
		return nil
	}
	var queue = make([]*TreeNode, 0)
	v := this.front()
	if *v == "null" {
		return nil
	}
	val, _ := strconv.Atoi(*v)
	root := &TreeNode{Val: val}
	queue = append(queue, root)
	for len(this.data) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head == nil {
			continue
		}
		if v := this.front(); v == nil {
			break
		} else if *v != "null" {
			val, _ = strconv.Atoi(*v)
			head.Left = &TreeNode{Val: val}
		}
		if v := this.front(); v == nil {
			break
		} else if *v != "null" {
			val, _ = strconv.Atoi(*v)
			head.Right = &TreeNode{Val: val}
		}
		queue = append(queue, head.Left, head.Right)
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
//leetcode submit region end(Prohibit modification and deletion)
