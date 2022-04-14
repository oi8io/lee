/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.66%)
 * Likes:    1610
 * Dislikes: 0
 * Total Accepted:    595.3K
 * Total Submissions: 920.3K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^4
 *
 *
 */

package problems

// @lc code=start
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	median := getMedian(nums)
	var i, j = 0, n - 2
	mid := 0
	for {
		if i > j {
			mid = i
			break
		}
		if nums[i] < median {
			i++
			continue
		}
		if nums[j] >= median {
			j--
			continue
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	nums[mid], nums[n-1] = nums[n-1], nums[mid]
	if n-mid == k {
		return median
	}
	// 右边的数字多于k个
	if n-mid > k {
		//fmt.Println("右边第", k, "个", nums[:mid], nums[mid], nums[mid+1:])
		return findKthLargest(nums[mid+1:], k)
	} else { //不够k个，则需要补k-(len(nums)-1-mid)
		//fmt.Println("左边第", k-(n-mid), "个", nums[:mid], nums[mid], nums[mid+1:])
		return findKthLargest(nums[:mid], k-(n-mid))
	}
}

func getMedian(nums []int) int {
	n := len(nums)
	var mid = len(nums) / 2
	if nums[0] > nums[mid] {
		nums[0], nums[mid] = nums[mid], nums[0]
	}
	if nums[0] > nums[n-1] {
		nums[0], nums[n-1] = nums[n-1], nums[0]
	}

	if nums[mid] < nums[n-1] {
		nums[mid], nums[n-1] = nums[n-1], nums[mid]
	}
	return nums[n-1]
}

// @lc code=end
