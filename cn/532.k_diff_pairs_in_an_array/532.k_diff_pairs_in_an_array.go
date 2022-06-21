//给你一个整数数组 nums 和一个整数 k，请你在数组中找出 不同的 k-diff 数对，并返回不同的 k-diff 数对 的数目。
//
// k-diff 数对定义为一个整数对 (nums[i], nums[j]) ，并满足下述全部条件：
//
//
// 0 <= i, j < nums.length
// i != j
// nums[i] - nums[j] == k
//
//
// 注意，|val| 表示 val 的绝对值。
//
//
//
// 示例 1：
//
//
//输入：nums = [3, 1, 4, 1, 5], k = 2
//输出：2
//解释：数组中有两个 2-diff 数对, (1, 3) 和 (3, 5)。
//尽管数组中有两个 1 ，但我们只应返回不同的数对的数量。
//
//
// 示例 2：
//
//
//输入：nums = [1, 2, 3, 4, 5], k = 1
//输出：4
//解释：数组中有四个 1-diff 数对, (1, 2), (2, 3), (3, 4) 和 (4, 5) 。
//
//
// 示例 3：
//
//
//输入：nums = [1, 3, 1, 5, 4], k = 0
//输出：1
//解释：数组中只有一个 0-diff 数对，(1, 1) 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -10⁷ <= nums[i] <= 10⁷
// 0 <= k <= 10⁷
//
// 👍 226 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findPairs(nums []int, k int) int {
	var numsMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}
	up := make(map[int]int)
	down := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := nums[i] - k
		if m, ok := numsMap[n]; ok && m != i {
			if _, ok := down[n]; !ok {
				up[nums[i]] = n
				down[n] = nums[i]
			}
		}
		n1 := nums[i] + k
		if m, ok := numsMap[n1]; ok && m != i {
			if _, ok := up[n1]; !ok {
				down[nums[i]] = n1
				up[n1] = nums[i]
			}
		}
	}
	return len(up)

}

//leetcode submit region end(Prohibit modification and deletion)
