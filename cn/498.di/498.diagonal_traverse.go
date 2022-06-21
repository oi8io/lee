//给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
//
//
//
// 示例 1：
//
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,4,7,5,3,6,8,9]
//
//
// 示例 2：
//
//
//输入：mat = [[1,2],[3,4]]
//输出：[1,2,3,4]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 10⁴
// 1 <= m * n <= 10⁴
// -10⁵ <= mat[i][j] <= 10⁵
//
// 👍 297 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findDiagonalOrder(mat [][]int) []int {
	//	向上 up   x+1 ,y+1
	//	项下 down x-1,y-1
	//	up转down x+1 || y-1
	//	down 转 up  x+1 || y-1
	var x, y int
	var up = true
	var anwser []int
	for {
		//最后一个点
		//fmt.Println(x, y, up)
		if x > len(mat)-1 || y > len(mat[0])-1 {
			break
		}
		anwser = append(anwser, mat[x][y])

		if up { //上有两种情况会被终止
			if y == len(mat[0])-1 { //最后一列,下一行继续开始
				x++
				up = !up
				continue
			}
			if x == 0 { //第一行，下一列继续开始
				y++
				up = !up
				continue
			}
			x--
			y++
		} else {
			if y == 0 { //  第一列，下一行
				if x == len(mat)-1 { // 最后一行，下一列
					y++
				} else {
					x++
				}
				up = !up
				continue
			}
			if x == len(mat)-1 { // 最后一行，下一列
				y++
				up = !up
				continue
			}
			y--
			x++
		}
	}
	return anwser
}

//leetcode submit region end(Prohibit modification and deletion)
