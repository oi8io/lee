//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
// 示例 1:
//
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
//
// 示例 2:
//
//
//输入: nums = [0]
//输出: [0]
//
//
//
// 提示:
//
//
//
// 1 <= nums.length <= 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能尽量减少完成的操作次数吗？
// 👍 1597 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	var count, start int
	var is_start bool
	//[0,1,0,3,12]
	//[1,1,0,3,12] start = 1 count = 2
	//start+count+str=i
	for i := 0; i <= len(nums); i++ {
		// 遇到0后判断之前是否有0未处理，如果未处理就先进行处理
		if i == len(nums) {
			if !is_start || i-start < count {
				break
			}
			for j := start + count; j < i; j++ {
				nums[j-count] = nums[j]
			}
			break
		}
		if nums[i] == 0 {
			count++
			if !is_start {
				is_start = true
				start = i //更新start的位置
				continue
			}
			if i-start < count { // 遇到连续0的情况
				continue
			}
			for j := start + count - 1; j < i; j++ {
				nums[j+1-count] = nums[j]
			}
			start = i + 1 - count
		}
	}
	for i := len(nums) - count; i < len(nums); i++ {
		nums[i] = 0
	}
}

//leetcode submit region end(Prohibit modification and deletion)
