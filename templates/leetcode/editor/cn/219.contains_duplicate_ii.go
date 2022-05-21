//给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i
//- j) <= k 。如果存在，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,1], k = 3
//输出：true
//
// 示例 2：
//
//
//输入：nums = [1,0,1,1], k = 1
//输出：true
//
// 示例 3：
//
//
//输入：nums = [1,2,3,1,2,3], k = 2
//输出：false
//
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// 0 <= k <= 10⁵
//
// 👍 481 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	k = k % len(nums)
	//滑动窗口
	var exist = make(map[int]struct{})

	for i := 0; i < k; i++ {
		if _, ok := exist[nums[i]]; ok {
			return true
		}
		exist[nums[i]] = struct{}{}
	}
	for i := k; i < len(nums); i++ {
		if _, ok := exist[nums[i]]; ok {
			return true
		}
		delete(exist, nums[i-k])
		exist[nums[i]] = struct{}{}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
