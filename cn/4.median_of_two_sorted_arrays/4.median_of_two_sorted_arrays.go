//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
// 👍 5530 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	var nums []int
	if n == 0 {
		return getMiddle(nums2)
	}
	if m == 0 {
		return getMiddle(nums1)
	}
	if n > 0 && nums1[n-1] < nums2[0] {
		nums = append(nums1, nums2...)
		return getMiddle(nums)
	} else if m > 0 && nums2[m-1] < nums1[0] {
		nums = append(nums2, nums1...)
		return getMiddle(nums)
	}
	for i, j := 0, 0; i < n && j < m; {
		if nums1[i] <= nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
		if i == n || j == m {
			if i <= n-1 {
				nums = append(nums, nums1[i:]...)
			} else if j <= m-1 {
				nums = append(nums, nums2[j:]...)
			}
		}

	}
	return getMiddle(nums)
}

func getMiddle(nums []int) float64 {
	n := len(nums)
	//fmt.Println(nums)
	if n%2 == 1 {
		return float64(nums[n/2])
	} else {
		return float64(nums[n/2-1]+nums[n/2]) / 2
	}
}

//leetcode submit region end(Prohibit modification and deletion)
