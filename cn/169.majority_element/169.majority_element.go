//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,3]
//输出：3
//
// 示例 2：
//
//
//输入：nums = [2,2,1,1,1,2,2]
//输出：2
//
//
//
//提示：
//
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
// 👍 1446 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
// 混战法。如果A同B不同。则A和B同时消失
func majorityElement(nums []int) int {
	var count = 1
	var king = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == king {
			count++
		} else {
			if count == 1 { //战死杀场
				king = nums[i]
			} else {
				count--
			}
		}
	}
	return king
}

// hash表法
func majorityElement1(nums []int) int {
	var xmap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		xmap[nums[i]]++
	}
	for i, i2 := range xmap {
		if i2 > len(nums)/2 {
			return i
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
