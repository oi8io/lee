package weekly

func equalPairs(grid [][]int) int {
	n := len(grid)
	var ans = 0
	for i := 0; i < n; i++ { // 行
		for j := 0; j < n; j++ { // 列
			k := 0
			for ; k < n; k++ {
				if grid[i][k] != grid[k][j] {
					break
				}
			}
			if k == n {
				ans++
			}
		}
	}
	return ans
}
