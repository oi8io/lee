package weekly

import (
	"sort"
)

func fillCups(amount []int) int {
	sort.Ints(amount)
	var cnt = 0
	for {
		if amount[1] == 0 {
			cnt += amount[2]
			break
		}
		cnt++
		amount[1]--
		amount[2]--
		sort.Ints(amount)
	}
	return cnt
}
