/**

 */

package weekly

func countPairs(n int, edges [][]int) int64 {
	parent, rank := initParent(n)
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		vertices := unionVertices(x, y, parent, rank)
		if vertices {
			//	do nothing
		}
	}
	var pm = make(map[int]int)
	for i := 0; i < len(parent); i++ {
		if parent[i] != -1 {
			r := findRoot(i, parent)
			pm[r]++
		} else {
			pm[i]++
		}
	}
	//fmt.Println(pm)
	if len(pm) == 1 {
		return 0
	}
	var answer int64
	for _, i2 := range pm {
		answer += int64((n - i2) * i2)
	}
	return answer / 2
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
func findChildren(n int, parent []int) int {
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
