//给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,0,1,1,1]
//输出：3
//解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
//
//
// 示例 2:
//
//
//输入：nums = [1,0,1,1,0,1]
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// nums[i] 不是 0 就是 1.
//
// 👍 318 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) int {
	var max = 0

	for i, j := 0, 0; j < len(nums) && i <= j; {
		if nums[j] == 0 { //重新开始
			if j-i > max {
				max = j - i
			}
			i = j
			i++
		}
		if j == len(nums)-1 {
			if j-i+1 > max {
				max = j - i + 1
			}
		}
		j++
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
