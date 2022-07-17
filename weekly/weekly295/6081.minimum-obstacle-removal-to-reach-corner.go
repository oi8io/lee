/**
6081. 到达角落需要移除障碍物的最小数目 显示英文描述
通过的用户数268
尝试过的用户数373
用户总通过次数274
用户总提交次数514
题目难度Hard
给你一个下标从 0 开始的二维整数数组 grid ，数组大小为 m x n 。每个单元格都是两个值之一：

0 表示一个 空 单元格，
1 表示一个可以移除的 障碍物 。
你可以向上、下、左、右移动，从一个空单元格移动到另一个空单元格。

现在你需要从左上角 (0, 0) 移动到右下角 (m - 1, n - 1) ，返回需要移除的障碍物的 最小 数目。



示例 1：



输入：grid = [[0,1,1],[1,1,0],[1,1,0]]
输出：2
解释：可以移除位于 (0, 1) 和 (0, 2) 的障碍物来创建从 (0, 0) 到 (2, 2) 的路径。
可以证明我们至少需要移除两个障碍物，所以返回 2 。
注意，可能存在其他方式来移除 2 个障碍物，创建出可行的路径。
示例 2：



输入：grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]]
输出：0
解释：不移除任何障碍物就能从 (0, 0) 到 (2, 4) ，所以返回 0 。


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 105
2 <= m * n <= 105
grid[i][j] 为 0 或 1
grid[0][0] == grid[m - 1][n - 1] == 0
https://leetcode.cn/contest/weekly-contest-295/problems/minimum-obstacle-removal-to-reach-corner/
*/

package weekly295

import (
	"fmt"
	"math"
)

func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	parseIdx := func(idx int) []int {
		return []int{idx / n, idx % n}
	}
	getIdx := func(x, y int) int {
		return x*n + y
	}

	f := make([][]int, m)
	g := make([]int, m*n)
	for i := m - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				f[i][j] = grid[i][j]
			} else {

				bottom := math.MaxInt
				if i+1 < m {
					bottom = f[i+1][j] + grid[i][j]
				}

				right := math.MaxInt
				if j+1 < n {
					right = f[i][j+1] + grid[i][j]
				}

				f[i][j] = min(bottom, right)
				g[getIdx(i, j)] = getIdx(i, j+1)
				if bottom < right {
					g[getIdx(i, j)] = getIdx(i+1, j)
				}
			}
		}
	}

	idx := getIdx(0, 0)
	for i := 1; i <= m+n; i++ {
		if i == m+n {
			continue
		}
		x := parseIdx(idx)[0]
		y := parseIdx(idx)[1]
		fmt.Println(x, y)
		idx = g[idx]
	}
	fmt.Println()
	return f[0][0]

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
