package weekly304

import (
	"sort"
)

func maximumGroups(grades []int) int {
	sort.Ints(grades)
	sum := make([]int, len(grades))
	sum[0] = grades[0]
	for i := 1; i < len(grades); i++ {
		sum[i] = sum[i-1] + grades[i]
	}
	//8,8
	var group = []int{sum[0]}
	var step, cur, before = 2, 0, sum[0]

	for {
		i := step + cur
		if i >= len(sum) {
			break
		}
		curSum := sum[i] - before
		if curSum > group[len(group)-1] {
			group = append(group, curSum)
			before += curSum
			cur = i
		}
		step++
	}
	return len(group)
}
