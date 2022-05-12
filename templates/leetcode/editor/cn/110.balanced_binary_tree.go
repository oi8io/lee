//<p>给定一个二叉树，判断它是否是高度平衡的二叉树。</p>
//
//<p>本题中，一棵高度平衡二叉树定义为：</p>
//
//<blockquote>
//<p>一个二叉树<em>每个节点 </em>的左右两个子树的高度差的绝对值不超过 1 。</p>
//</blockquote>
//
//<p> </p>
//
//<p><strong>示例 1：</strong></p>
//<img alt="" src="https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg" style="width: 342px; height: 221px;" />
//<pre>
//<strong>输入：</strong>root = [3,9,20,null,null,15,7]
//<strong>输出：</strong>true
//</pre>
//
//<p><strong>示例 2：</strong></p>
//<img alt="" src="https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg" style="width: 452px; height: 301px;" />
//<pre>
//<strong>输入：</strong>root = [1,2,2,3,3,null,null,4,4]
//<strong>输出：</strong>false
//</pre>
//
//<p><strong>示例 3：</strong></p>
//
//<pre>
//<strong>输入：</strong>root = []
//<strong>输出：</strong>true
//</pre>
//
//<p> </p>
//
//<p><strong>提示：</strong></p>
//
//<ul>
//	<li>树中的节点数在范围 <code>[0, 5000]</code> 内</li>
//	<li><code>-10<sup>4</sup> <= Node.val <= 10<sup>4</sup></code></li>
//</ul>
//<div><li>👍 995</li><li>👎 0</li></div>

package problems

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, b := GetHeight(root)
	return !b
}

func GetHeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 1, false
	}
	var leftHeight, rightHeight = 0, 0
	var b bool
	if root.Left != nil {
		leftHeight, b = GetHeight(root.Left)
		if b {
			return leftHeight, b
		}
	} else {
		leftHeight = 1
	}
	if root.Right != nil {
		rightHeight, b = GetHeight(root.Right)
		if b {
			return leftHeight, b
		}
	} else {
		rightHeight = 1
	}
	if abs(leftHeight, rightHeight) > 1 {
		return max(leftHeight, rightHeight) + 1, true
	} else {
		return max(leftHeight, rightHeight) + 1, false
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
