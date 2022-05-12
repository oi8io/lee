//给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次
//的整数，并以数组形式返回。
//
// 你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[2,3]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,2]
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= n
// nums 中的每个元素出现 一次 或 两次
//
// 👍 504 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
//审题，范围在1到n之间
func findDuplicates1(nums []int) []int {
	var ret []int
	//fmt.Println(nums)
	var abs = func(num int) int {
		if num < 0 {
			return 0 - num
		}
		return num
	}
	//两次更新数字，如果可以采用 负负为正 或者 +n操作
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		index := abs(num) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		} else {
			ret = append(ret, index+1)
		}
		//fmt.Println(num, index, nums)
	}
	return ret
}

func findDuplicates(nums []int) []int {
	var ret []int
	n := len(nums) + 1
	//两次更新数字，如果可以采用 负负为正 或者 +n操作
	for i := 0; i < len(nums); i++ {
		num := nums[i] % n
		index := num - 1
		if nums[index] < n { //未更新过
			nums[index] = n + nums[index]
		} else { //已经更新过一次了
			ret = append(ret, index+1)
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
