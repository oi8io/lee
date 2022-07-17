package biweekly82

import (
	"math/rand"
	"sort"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = abs(nums2[i], nums1[i])
	}
	sort.Ints(a)

	return 0
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func shuffle(arr sort.IntSlice) []int {
	for i := arr.Len(); i >= 1; i-- {
		r := rand.Intn(i) //左闭右开
		arr.Swap(i-1, r)
	}
	return arr
}
