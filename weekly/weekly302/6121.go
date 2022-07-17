package weekly

import "sort"

func maximumSum(nums []int) int {
	sort.Ints(nums)
	var m = make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		s := sum(nums[i])
		m[s] = append(m[s], nums[i])
	}
	var ans = -1
	for _, arr := range m {
		if len(arr) == 1 {
			continue
		}
		x := arr[len(arr)-2] + arr[len(arr)-1]
		if x > ans {
			ans = x
		}
	}
	return ans
}

func sum(n int) int {
	var r = 0
	for n > 0 {
		r += n % 10
		n = n / 10
	}
	return r
}
