package cn

import (
	"fmt"
	"strconv"
	"strings"
)

func GetMatrix(m, n int) [][]int {
	var matrix = make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = i*n + j + 1
		}
	}
	return matrix
}

func PrintMatrix(m [][]int) {
	s := make([]string, 0)
	for i := 0; i < len(m); i++ {
		x := make([]string, 0)
		for j := 0; j < len(m[0]); j++ {
			x = append(x, strconv.Itoa(m[i][j]))
		}
		join := fmt.Sprintf("[%s]", strings.Join(x, ","))
		s = append(s, join)
	}
	join := fmt.Sprintf("[%s]", strings.Join(s, ","))
	fmt.Println(join)
}
func PrintArray(m []int) {
	x := make([]string, 0)
	for j := 0; j < len(m); j++ {
		x = append(x, strconv.Itoa(m[j]))
	}
	join := fmt.Sprintf("[%s]", strings.Join(x, ","))
	fmt.Println(join)
}
