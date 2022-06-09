//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
//
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]
//
//
// 示例 2:
//
//
//输入: nums = [-1,1,0,-3,3]
//输出: [0,0,9,0,0]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内
//
//
//
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
// 👍 1162 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	var answer = make([]int, len(nums))
	n := len(nums) - 1
	left, right := 1, 1
	for i := 0; i <= n; i++ {
		if i <= n-i {
			answer[i] = 1
			answer[n-i] = 1
		}
		answer[i] *= left
		left *= nums[i]
		answer[n-i] *= right
		right *= nums[n-i]
	}
	return answer
}
func productExceptSelf1(nums []int) []int {
	var left, right []int
	n := len(nums) - 1
	for i := 0; i <= n; i++ {
		beforeLeft, beforeRight := 1, 1
		if i > 0 {
			beforeLeft = left[i-1]
			beforeRight = right[i-1]
		}
		left = append(left, beforeLeft*nums[i])
		right = append(right, beforeRight*nums[n-i])
	}
	//fmt.Println(nums)
	//fmt.Println(left, right)
	var answer []int
	for i := 0; i <= n; i++ {
		if i == 0 {
			answer = append(answer, right[n-1])
		} else if i == len(nums)-1 {
			answer = append(answer, left[n-1])
		} else {
			//-1 -1 0 0 0  , 0 0 0 -9 3
			answer = append(answer, left[i-1]*right[n-i-1])
		}
	}
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
