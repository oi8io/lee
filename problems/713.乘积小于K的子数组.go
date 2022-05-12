/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于K的子数组
 *
 * https://leetcode-cn.com/problems/subarray-product-less-than-k/description/
 *
 * algorithms
 * Medium (43.93%)
 * Likes:    444
 * Dislikes: 0
 * Total Accepted:    49.5K
 * Total Submissions: 107.9K
 * Testcase Example:  '[10,5,2,6]\n100'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,5,2,6], k = 100
 * 输出：8
 * 解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2]、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
 * 需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3], k = 0
 * 输出：0
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * 1 <= nums[i] <= 1000
 * 0 <= k <= 10^6
 *
 *
 */

package problems

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {

	var sum, ret = 1, 0
	for n := 0; n < len(nums); n++ { //每次选几个数字
		sum = sum * nums[n]
		if sum == 1 && k > 1 {
			ret += len(nums) - n
			continue
		}
		tmpSum := sum
		prev := 0
		newRet := 0
		for i := 0; i < len(nums)-n; i++ {
			start, end := i, (i+n)%len(nums)+1
			var arr []int
			if start >= end {
				break
				// arr = append(append([]int{}, nums[start:]...), nums[:end]...)
			} else {
				arr = nums[start:end]
			}
			if prev != 0 && len(arr) > 0 {
				tmpSum = tmpSum / prev * arr[len(arr)-1]
			}
			if tmpSum < k {
				newRet++
			}
			prev = nums[start]
			if start == end {
				break
			}
			// fmt.Printf("[%d,%d] array: %+v sum: %d %+v\n", n, i, arr, tmpSum, tmpSum < k)
		}
		if newRet == 0 {
			break
		}
		ret += newRet
		//fmt.Println()
	}
	return ret
}

// @lc code=end
