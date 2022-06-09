//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长
//度。如果不存在符合条件的子数组，返回 0 。
//
//
//
// 示例 1：
//
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
//
//
// 示例 2：
//
//
//输入：target = 4, nums = [1,4,4]
//输出：1
//
//
// 示例 3：
//
//
//输入：target = 11, nums = [1,1,1,1,1,1,1,1]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
//
//
// 进阶：
//
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
//
// 👍 1165 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	//[2,3,1,2,4,3]
	//	 start -> 0 end ->3 val = 8 new=6,start = 3
	var i, j, min, sum, has = 0, 0, len(nums), nums[0], false

	for {
		if i == len(nums) || j > len(nums)-1 {
			break
		}
		//fmt.Println(i, j, sum)
		if sum < target { // i 向右移动
			if i+j < len(nums)-1 {
				sum += nums[i+j+1]
				j++
			} else {
				break
			}
		} else {
			has = true
			if j+1 < min {
				min = j + 1
			}
			sum -= nums[i]
			i++
			j--
		}
	}
	if !has {
		return 0
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
