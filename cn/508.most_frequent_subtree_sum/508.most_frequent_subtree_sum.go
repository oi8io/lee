//给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
//
// 一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
//
//
//
// 示例 1：
//
//
//
//
//输入: root = [5,2,-3]
//输出: [2,-3,4]
//
//
// 示例 2：
//
//
//
//
//输入: root = [5,2,-5]
//输出: [2]
//
//
//
//
// 提示:
//
//
// 节点数在 [1, 10⁴] 范围内
// -10⁵ <= Node.val <= 10⁵
//
// 👍 173 👎 0

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
func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	sumMap := make(map[int]int)
	rootSum := getFrequentTreeSum(root, sumMap)
	var maxCount = sumMap[rootSum]
	var sumGroup = make(map[int][]int)
	for sum, count := range sumMap {
		if count > maxCount {
			maxCount = count
		}
		sumGroup[count] = append(sumGroup[count], sum)
	}
	return sumGroup[maxCount]
}

func getFrequentTreeSum(root *TreeNode, m map[int]int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		m[root.Val]++
		return root.Val
	}
	leftSum := getFrequentTreeSum(root.Left, m)
	rightSum := getFrequentTreeSum(root.Right, m)
	sum := leftSum + rightSum + root.Val
	m[sum]++
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
