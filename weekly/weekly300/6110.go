package weekly

import . "github.com/oi8io/lee/cn/common"

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	var matrix = make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	h := head
	var x, y = 0, 0
	var cnt = 0
	var visited = make(map[int]bool)
	for h != nil {
		i, j := getNext(m, n, x, y, visited)
		matrix[i][j] = h.Val
		visited[i*m+j] = true
		x, y = i, j

		cnt++
	}
	return matrix
}
func getNext(m, n, x, y int, visited map[int]bool) (int, int) {
	if x <= m/2 { //正着走 左到右 上到下
		//向右
		if y > n/2 { //向下

		}
	} else {

	}
	return 0, 0
}
