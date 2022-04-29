package problems

import (
	"fmt"
)

//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[7,4,1],[8,5,2],[9,6,3]]
//
//
// 示例 2：
//
//
//输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
//
//
//
//
// 提示：
//
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000
//
//
//
// Related Topics 数组 数学 矩阵 👍 1219 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	var (
		n = len(matrix)
	)

	/*	for _, ints := range matrix {
		for _, i2 := range ints {
			fmt.Printf("%02d ", i2)
		}
		fmt.Println()
	}*/
	//i表示行数，旋转之后变成对应的列，刚好得到y=n-1-i
	//j表示列数，旋转之后第一列对应第一行，则x=j
	getNext := func(i, j int) (int, int) {
		y := n - 1 - i
		x := j
		return x, y
	}

	// solveNQueensAnswers 表示由外向内第几个圈
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ { // j的范围就是从i到n-1-i
			x, y := i, j
			tmp := matrix[x][y]
			var next = 0
			for i := 0; i < 4; i++ { //每次转一圈把四个相关的元素替换掉
				nx, ny := getNext(x, y)
				next = matrix[nx][ny]
				matrix[nx][ny] = tmp
				tmp = next
				//fmt.Printf("(%d,%d)->(%d,%d),[%d]->[%d] \n", solveNQueensAnswers, y, nx, ny, matrix[nx][ny], tmp)
				x, y = nx, ny
			}
			//fmt.Println("----------------")
		}
		//fmt.Println("++++++++++++")

	}
	for _, ints := range matrix {
		for _, i2 := range ints {
			fmt.Printf("%d ", i2)
		}
		//fmt.Println()
	}
}

/**
11 -> 41
21 -> 42
31 -> 43
41 -> 44

00->30
01->31
11->12


13->30




22->23
23->33
33->32
32->22
solveNQueensAnswers = n-y


11->22  23 04


*/
func rotate1(matrix [][]int) [][]int {
	for _, ints := range matrix {
		for _, i2 := range ints {
			fmt.Printf("%02d ", i2)
		}
		fmt.Println()
	}
	n := len(matrix)
	fmt.Println(n)
	if n == 1 {
		return matrix
	}
	var res = make([][]int, n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			y := n - 1 - i
			x := j
			if res[x] == nil {
				res[x] = make([]int, n)
			}
			res[x][y] = matrix[i][j]
		}
		fmt.Println()
	}
	for _, ints := range res {
		for _, i2 := range ints {
			fmt.Printf("%02d ", i2)
		}
		fmt.Println()
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
