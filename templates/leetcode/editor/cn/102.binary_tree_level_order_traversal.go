//<p>给你二叉树的根节点 <code>root</code> ，返回其节点值的 <strong>层序遍历</strong> 。 （即逐层地，从左到右访问所有节点）。</p>
//
//<p>&nbsp;</p>
//
//<p><strong>示例 1：</strong></p>
//<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" style="width: 277px; height: 302px;" />
//<pre>
//<strong>输入：</strong>root = [3,9,20,null,null,15,7]
//<strong>输出：</strong>[[3],[9,20],[15,7]]
//</pre>
//
//<p><strong>示例 2：</strong></p>
//
//<pre>
//<strong>输入：</strong>root = [1]
//<strong>输出：</strong>[[1]]
//</pre>
//
//<p><strong>示例 3：</strong></p>
//
//<pre>
//<strong>输入：</strong>root = []
//<strong>输出：</strong>[]
//</pre>
//
//<p>&nbsp;</p>
//
//<p><strong>提示：</strong></p>
//
//<ul>
//	<li>树中节点数目在范围 <code>[0, 2000]</code> 内</li>
//	<li><code>-1000 &lt;= Node.val &lt;= 1000</code></li>
//</ul>
//<div><li>👍 1311</li><li>👎 0</li></div>

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret = make([][]int, 0)
	ret = level(root, 0, ret)
	return ret
}

func level(root *TreeNode, l int, ret [][]int) [][]int {
	if root == nil {
		return ret
	}
	if len(ret) < l+1 {
		ret = append(ret, make([]int, 0))
	}
	ret[l] = append(ret[l], root.Val)
	ret = level(root.Left, l+1, ret)
	ret = level(root.Right, l+1, ret)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
