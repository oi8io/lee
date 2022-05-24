//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否
//则返回 -1。
//
//
//示例 1:
//
// 输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//
// 示例 2:
//
// 输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//
//
//
// 提示：
//
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
//
// 👍 832 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	var min, max = 0, len(nums) - 1
	var mid int
	for {
		//fmt.Println(min, max, mid)
		if min >= max-1 {
			break
		}
		mid = (min + max) / 2
		if nums[mid] > target {
			max = mid
		} else if nums[mid] == target {
			return mid
		} else {
			if mid == min {
				min++
			} else {
				min = mid
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
