package questions

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
//
//
// 示例 1：
//
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：1
//
//
// 示例 3：
//
//
//输入：nums = [5,4,-1,7,8]
//输出：23
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
// Related Topics 数组 分治 动态规划 👍 4479 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	var max, current int
	var is1 = true

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if is1 {
			is1 = false
			max = nums[i]
			current = nums[i]
			continue
		}
		if num < 0 { // 小于0 则计算左右
			if current+num > 0 {
				current = current + num
			} else { // 截断
				current = num
			}
		} else {
			if current < 0 {
				current = num
			} else {
				current = current + num
			}
		}
		if max < current {
			max = current
		}
	}
	return max
}

/**
-2, 1, -3, 4, -1, 2, 1, -5, 4

1 大于 -2 且 -2<0 1,0

1>-3 且-3<0 且1-3<0 故截断 保留 1,0
4
4+-1=3 > 0 4, 要想与后面的相加，则用 4,3
3+2>4 故 5,5

5+1 = 6> 5 ,  6 6

6+-5 = 1,  6  1

1+4 = 5 <6  6,5

*/

//func subArray(nums []int) (bool, int) {
//
//}

//leetcode submit region end(Prohibit modification and deletion)
