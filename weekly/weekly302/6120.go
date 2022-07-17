package weekly

func numberOfPairs(nums []int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var ans = make([]int, 2)
	for _, i2 := range m {
		ans[0] += i2 / 2
		ans[1] += i2 % 2
	}
	return ans
}
