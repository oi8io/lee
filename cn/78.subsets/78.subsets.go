//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
// 👍 1708 👎 0

package cn

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var x = make([][]int, 1)
	x[0] = make([]int, 0)
	for j := 1; j <= n; j++ {
		for i := 0; i < n; i++ {
			t := try(nums, n, i, j, []int{})
			x = append(x, t...)
		}
	}
	return x
}

func try(nums []int, n, i, target int, arr []int) [][]int {
	arr = copyS(arr)
	arr = append(arr, nums[i])
	if len(arr) == target {
		return [][]int{arr}
	}
	var ret = make([][]int, 0)
	for k := i + 1; k < len(nums); k++ {
		t := try(nums, n, k, target, arr)
		ret = append(ret, t...)
	}
	return ret
}

func copyS(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
