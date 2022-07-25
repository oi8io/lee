package weekly

import (
	"math/bits"
)

func countExcellentPairs(nums []int, k int) int64 {
	var m, e = make(map[int]int), make(map[int]bool)
	for _, n := range nums {
		if !e[n] {
			count := bits.OnesCount(uint(n))
			m[count]++
			e[n] = true
		}
	}
	var ans int64 = 0
	for i, i2 := range m {
		for j, j2 := range m {
			if i+j >= k {
				ans += int64(i2 * j2)
			}
		}
	}
	return ans
}
