//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方
//式重构得到原数据。
//
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串
//反序列化为原始的树结构。
//
// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的
//方法解决这个问题。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,null,4,5]
//输出：[1,2,3,null,null,4,5]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中结点数在范围 [0, 10⁴] 内
// -1000 <= Node.val <= 1000
//
// 👍 903 👎 0

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

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

type Codec struct {
	data []string
}

func Constructor() Codec {
	return Codec{make([]string, 0)}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return c.serializeLevel(root)
}

// Serializes a tree to a single string.
func (c *Codec) serializeLevel(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head == nil {
			c.data = append(c.data, "null")
		} else {
			v := strconv.Itoa(head.Val)
			c.data = append(c.data, v)
			queue = append(queue, head.Left)
			queue = append(queue, head.Right)
		}
	}
	i := len(c.data) - 1
	for ; i >= 0; i-- {
		if c.data[i] != "null" {
			break
		}
	}
	c.data = c.data[:i+1]
	return strings.Join(c.data, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" || data == "null" {
		return nil
	}
	return c.deserializeLevel(data)
}
func (c *Codec) deserializeLevel(data string) *TreeNode {
	c.data = strings.Split(data, ",")
	v, _ := strconv.Atoi(c.data[0])
	root := &TreeNode{Val: v}
	queue := []*TreeNode{root}
	c.data = c.data[1:]
	for len(queue) > 0 {
		if len(c.data) == 0 {
			break
		}
		head := queue[0]
		queue = queue[1:]
		if c.data[0] != "null" {
			v, _ := strconv.Atoi(c.data[0])
			head.Left = &TreeNode{Val: v}
			queue = append(queue, head.Left)
		}
		c.data = c.data[1:]
		if len(c.data) == 0 {
			break
		}
		if c.data[0] != "null" {
			v, _ := strconv.Atoi(c.data[0])
			head.Right = &TreeNode{Val: v}
			queue = append(queue, head.Right)
		}
		c.data = c.data[1:]
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)
