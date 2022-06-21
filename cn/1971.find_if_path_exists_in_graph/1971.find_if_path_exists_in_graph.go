//有一个具有 n个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges 表示，其中
//edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。
//
// 请你确定是否存在从顶点 start 开始，到顶点 end 结束的 有效路径 。
//
// 给你数组 edges 和整数 n、start和end，如果从 start 到 end 存在 有效路径 ，则返回 true，否则返回 false 。
//
//
//
// 示例 1：
//
//
//输入：n = 3, edges = [[0,1],[1,2],[2,0]], start = 0, end = 2
//输出：true
//解释：存在由顶点 0 到顶点 2 的路径:
//- 0 → 1 → 2
//- 0 → 2
//
//
// 示例 2：
//
//
//输入：n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], start = 0, end = 5
//输出：false
//解释：不存在由顶点 0 到顶点 5 的路径.
//
//
//
//
// 提示:
//
//
// 1 <= n <= 2 * 10⁵
// 0 <= edges.length <= 2 * 10⁵
// edges[i].length == 2
// 0 <= ui, vi <= n - 1
// ui != vi
// 0 <= start, end <= n - 1
// 不存在双向边
// 不存在指向顶点自身的边
//
// 👍 36 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func validPath(n int, edges [][]int, source int, destination int) bool {
	parent, rank := initParent(n)
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		vertices := unionVertices(x, y, parent, rank)
		if vertices {
			//	do nothing
		}
	}
	sr := findRoot(source, parent)
	dr := findRoot(destination, parent)
	return sr == dr
}

func initParent(n int) (parent []int, rank []int) {
	parent = make([]int, n, n)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}
	rank = make([]int, n, n)
	return parent, rank
}

func findRoot(n int, parent []int) int {
	for parent[n] != -1 {
		n = parent[n]
	}
	return n
}

// 并查集
func unionVertices(x, y int, parent []int, rank []int) bool {
	xRoot := findRoot(x, parent)
	yRoot := findRoot(y, parent)
	if xRoot == yRoot {
		return false
	}
	if rank[xRoot] > rank[yRoot] {
		parent[yRoot] = xRoot
	} else if rank[xRoot] < rank[yRoot] {
		parent[yRoot] = xRoot
	} else {
		rank[yRoot]++
		parent[yRoot] = xRoot
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
