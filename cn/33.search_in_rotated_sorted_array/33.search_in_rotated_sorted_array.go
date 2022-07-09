//整数数组 nums 按升序排列，数组中的值 互不相同 。
//
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[
//k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2
//,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//
//
// 示例 2：
//
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//
// 示例 3：
//
//
//输入：nums = [1], target = 0
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// nums 中的每个值都 独一无二
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// -10⁴ <= target <= 10⁴
//
// 👍 2175 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	//	需不需要找到分割点
	n := len(nums)
	lo, hi, m := 0, n, 0
	for {
		m = lo + (hi-lo)>>1
		if lo >= hi-1 {
			break
		}
		if nums[m] == target {
			return m
		}
		if nums[m] > nums[0] { //上升
			if target < nums[0] || target > nums[m] {
				lo = m
			} else {
				hi = m
			}
		} else { //下降，进入旋转
			if target > nums[n-1] || target < nums[m] { //
				hi = m
			} else {
				lo = m
			}
		}
	}
	if nums[m] == target {
		return m
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
