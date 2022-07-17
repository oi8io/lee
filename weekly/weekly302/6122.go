package weekly

import (
	"sort"
)

type Val struct {
	k string
	v int
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	n, m := len(nums), len(nums[0])

	var xs = make(map[int][]Val)
	var ans = make([]int, len(queries))
	for k := 0; k < len(queries); k++ {
		start := m - queries[k][1]
		x := queries[k][0] - 1
		if ns, ok := xs[start]; ok {
			ans[k] = ns[x].v
			continue
		}
		var ks = make([]Val, n)
		for i := 0; i < n; i++ {
			ks[i] = Val{nums[i][start:], i}
		}
		sort.Slice(ks, func(i, j int) bool {
			if ks[i].k == ks[j].k {
				return ks[i].v < ks[j].v
			} else {
				return ks[i].k < ks[j].k
			}
		})

		xs[start] = ks
		ans[k] = xs[start][x].v
	}
	return ans
}
