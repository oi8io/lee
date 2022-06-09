//给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
//
//
//
// 示例 1：
//
//
//输入：nums = [-10,-3,0,5,9]
//输出：[0,-3,9,-10,null,5]
//解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
//
//
//
// 示例 2：
//
//
//输入：nums = [1,3]
//输出：[3,1]
//解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 按 严格递增 顺序排列
//
// 👍 1035 👎 0

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
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	root := sortedArrayToBST1(nums, 0, n-1)
	return root
}

func sortedArrayToBST1(nums []int, start, end int) *TreeNode {
	if start >= end {
		return &TreeNode{Val: nums[start]}
	}
	mid := (start + end) / 2

	root := &TreeNode{Val: nums[mid]}
	if start <= mid-1 {
		root.Left = sortedArrayToBST1(nums, start, mid-1)
	}
	if mid+1 <= end {
		root.Right = sortedArrayToBST1(nums, mid+1, end)
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
