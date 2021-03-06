package problems

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
// Related Topics 数组 回溯 👍 1878 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

/**
解题思路
1. 两个相同的集合中连线，相同的不连接，线条总数相加
2. 分解子问题，1，2，3 分别将分为 【1】，【2，3】，【2】，【1，3】，【3】，【1，2】子问题，【1，2】同样 【1】，【2】当只有一个数字的时候返回本身的数组，


船夫过河问题：船夫要把一匹狼、一只羊和一棵白菜运过河。只要船夫不在场，羊就会吃白菜、狼就会吃羊。船夫的船每次只能运送一种东西。怎样把所有东西都运过河？这是线性规划的问题。


第一次，先将羊运送到对面
第二次，将白菜运送至对面,把羊拉回来
第三次，将狼运过去
第四次，将羊也运过去

*/
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	n := len(nums)
	var current [][]int
	for i := 0; i < n; i++ {
		var nnums []int
		nnums = append(nnums, nums[0:i]...)
		nnums = append(nnums, nums[i+1:]...)
		rows := permute(nnums)
		for _, row := range rows {
			row = append(row, nums[i])
			current = append(current, row)
		}
	}
	return current
}

//leetcode submit region end(Prohibit modification and deletion)
