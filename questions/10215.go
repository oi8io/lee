package questions

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
//
// 👍 1585 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		max := nums[0]
		var p = 0
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] >= max {
				max = nums[j]
				p = j
			}
		}
		if p != len(nums)-i-1 {
			nums[p], nums[len(nums)-1-i] = swap(nums[p], nums[len(nums)-1-i])
		}
	}
	return nums[len(nums)-k]
}
func swap(a, b int) (int, int) {
	return b, a
}

//leetcode submit region end(Prohibit modification and deletion)
