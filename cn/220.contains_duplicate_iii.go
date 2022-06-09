//给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <=
//t ，同时又满足 abs(i - j) <= k 。
//
// 如果存在则返回 true，不存在返回 false。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,1], k = 3, t = 0
//输出：true
//
// 示例 2：
//
//
//输入：nums = [1,0,1,1], k = 1, t = 2
//输出：true
//
// 示例 3：
//
//
//输入：nums = [1,5,9,1,5,9], k = 2, t = 3
//输出：false
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 2 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// 0 <= k <= 10⁴
// 0 <= t <= 2³¹ - 1
//
// 👍 607 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var tree *TreeNode
	for i := 0; i < len(nums); i++ {
		bst1 := searchBST220(tree, nums[i], t)
		if bst1 != nil {
			return true
		}
		n := &TreeNode{Val: nums[i]}
		tree = insertIntoBST220(tree, n)
		if i-k >= 0 { //需要删除元素
			tree = deleteBSTNode220(tree, nums[i-k])
		}
	}
	return false
}

func searchBST220(root *TreeNode, val, t int) *TreeNode {
	if root == nil {
		return nil
	}
	if absSub(root.Val, val) <= t {
		return root
	}
	if root.Val < val {
		return searchBST220(root.Right, val, t)
	}
	if root.Val > val {
		return searchBST220(root.Left, val, t)
	}
	return nil
}

func absSub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func insertIntoBST220(root *TreeNode, node *TreeNode) *TreeNode {
	if root == nil {
		return node
	}
	if root.Val < node.Val {
		root.Right = insertIntoBST220(root.Right, node)
	}
	if root.Val > node.Val {
		root.Left = insertIntoBST220(root.Left, node)
	}
	return root
}

func deleteBSTNode220(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key > root.Val && root.Right != nil {
		root.Right = deleteBSTNode220(root.Right, key)
		return root
	} else if key < root.Val && root.Left != nil {
		root.Left = deleteBSTNode220(root.Left, key)
		return root
	} else if key == root.Val {
		if root.Right == nil { // 右子树为空 0/1个子树
			return root.Left
		}
		if root.Left == nil { // 左子树为空 0/1个子树
			return root.Right
		}
		//两个孩子的情况，找右子节点 todo 后继节点
		min := FindBSTMin220(root.Right)
		root.Val = min.Val
		root.Right = deleteBSTNode220(root.Right, min.Val)
		return root
	}
	return root
}

func FindBSTMin220(root *TreeNode) *TreeNode {
	for root != nil && root.Left != nil {
		root = root.Left
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
