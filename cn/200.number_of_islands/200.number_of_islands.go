//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
// 👍 1794 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	return numIslands2(grid)
}
func numIslands2(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var w, l byte = '0', '1'
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= m || y >= n {
			return
		}
		if grid[x][y] == l {
			grid[x][y] = w
			dfs(x+1, y)
			dfs(x-1, y)
			dfs(x, y-1)
			dfs(x, y+1)
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == l {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
func numIslands1(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var w byte = '0'

	var ans = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == w {
				continue
			}
			if i == 0 || grid[i-1][j] == w { // top with zero
				if j == 0 || grid[i][j-1] == w { // left with zero
					ans++
				}
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
