//给你一个整数数组 nums，请你将该数组升序排列。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// -5 * 10⁴ <= nums[i] <= 5 * 10⁴
//
// 👍 604 👎 0

package cn

import (
	"container/heap"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	var ret []int
	//ret = merge(nums)
	//ret = bubble(nums)
	ret = insertSort(nums)
	//ret = heapSort(nums)
	return ret
}

// 快速排序
func qsort(nums []int) []int {
	return nil
}

func shell(nums []int) []int {
	return nil
}

type mh struct {
	sort.IntSlice
}

func (m *mh) Push(x interface{}) {}
func (m *mh) Pop() interface{} {
	x := m.IntSlice[0]
	m.IntSlice[0] = m.IntSlice[m.Len()-1]
	m.IntSlice = m.IntSlice[:m.Len()-1]
	heap.Fix(m, 0)
	return x
}
func heapSort(nums []int) []int {
	var h = &mh{nums}
	heap.Init(h)
	var ans = make([]int, 0)
	for h.Len() > 0 {
		ans = append(ans, h.Pop().(int))
	}
	return ans
}

func merge(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}
	// 拆分成左右两部分
	ret1 := merge(nums[:n/2])
	ret2 := merge(nums[n/2:])
	out := merge2(ret1, ret2)
	//fmt.Println(nums, out, len(nums)/2, nums[:len(nums)/2], nums[len(nums)/2:], ret1, ret2)
	return out
}

func merge2(a1 []int, a2 []int) (out []int) {
	j, i, m, n := 0, 0, len(a1), len(a2)
	// 循环，避免越界
	for i <= m && j <= n {
		if i == m { // 左边的数组遍历完成，把所有的右边剩余部分追加到数组
			out = append(out, a2[j:]...)
			break
		}
		if j == n { // 同上
			out = append(out, a1[i:]...)
			break
		}
		// 将较小的数字追加到数组
		if a1[i] < a2[j] {
			out = append(out, a1[i])
			i++
		} else {
			out = append(out, a2[j])
			j++
		}
	}
	return out
}

// 插入排序
func insertSort(nums sort.IntSlice) []int {
	nums1 := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		insert(nums, i, nums1)
		//for j := i; j >= 1; j-- {
		//	if nums.Less(j, j-1) {
		//		nums.Swap(j, j-1)
		//	}
		//}
	}
	return nums
}

func insert(nums sort.IntSlice, i int, nums1 []int) {
	if nums[i] >= nums[i-1] {
		return
	}
	lo, m, hi := 0, 0, i
	for {
		m = lo + (hi-lo)>>1
		if nums[m] == nums[i] {
			break
		}
		if lo >= hi {
			break
		}
		if nums[m] > nums[i] {
			hi = m
		} else {
			if lo == m {
				lo++
			} else {
				lo = m
			}
		}
	}

	x := nums[i]
	copy(nums1, nums[m:i])
	copy(nums[m+1:i+1], nums1[:i-m+1])
	nums[m] = x
}

// 冒泡排序
func bubble(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// 系统快排
func system(nums []int) []int {
	sort.Ints(nums)
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
