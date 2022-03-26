package pk

import (
	"fmt"
)

/**

 */

type stack struct {
	data   []int
	offset int
	len    int
	name   string
}

func (s *stack) pop() int {
	ret := s.data[s.offset-1]
	s.offset = s.offset - 1
	s.data = s.data[0:s.offset]
	s.len--
	return ret
}

func (s *stack) getLast() int {
	return s.data[s.offset-1]
}

func (s *stack) push(i int) {
	s.offset = s.offset + 1
	s.data = append(s.data, i)
	s.len++
}

func PushAble(from, to *stack) bool {
	if from.len == 0 {
		return false
	}

	if to.len == 0 {
		return true
	}
	if from.data[from.offset-1] < to.data[to.offset-1] {
		return true
	}
	return false
}

func Hanoi(n int) int {
	a, b, c := &stack{name: "From"}, &stack{name: "Mid"}, &stack{name: "To"}
	for i := n; i > 0; i-- {
		a.push(i)
	}
	//fmt.Println(a)
	HanoiTower(a, b, c)
	return 0
}
func HanoiTower(from, mid, to *stack) int {
	// finish
	if from.len == 0 {
		return 0
	}
	fmt.Println(from, mid, to)
	// 如果目标塔不为空且当前值大于目标
	if to.len > 0 && from.data[from.offset-1] > to.data[to.offset-1] {
		HanoiTower(to, from, mid)
	}
	num := from.pop()
	to.push(num)

	if mid.len > 0 {
		HanoiTower(mid, from, to)
	}

	fmt.Println(from, mid, to)
	// 没有传输完成，继续传输
	if from.len >= 0 && from.getLast() > num {
		HanoiTower(from, mid, to)
	}

	return 0
}

// TowerOfHanoi 分解子问题，先将n-1挪到b的位置，再将n挪到n的c的位置，再将n-1由b挪到c
func TowerOfHanoi(n int, a, b, c string) {
	if n == 0 {
		return
	}
	TowerOfHanoi(n-1, a, c, b)
	fmt.Printf("num %d from [%s] TO [%s]\n", n, a, c)
	TowerOfHanoi(n-1, b, a, c)
}

type tower struct {
	data []int
	name string
}

// TowerOfHanoi2 分解子问题，先将n-1挪到b的位置，再将n挪到n的c的位置，再将n-1由b挪到c
func TowerOfHanoi2(a, b, c []int) {
	if len(a) == 0 {
		return
	}
	TowerOfHanoi2(a[1:], c, b)
	c = append(c, a[0])
	if len(a) > 1 {
		a = a[1:]
	}
	fmt.Printf("num %d from [%v] TO [%v]\n", a[0], a, c)
	TowerOfHanoi2(b, a, c)
	fmt.Println(a, b, c)
}

func Hanota(A, B, C []int) []int {
	fmt.Println(A, B, C)
	if len(A) > 0 {
		Hanota(A[1:], C, B)
		C = append(C, A[0])
		Hanota(B, []int{}, C)
	}
	fmt.Println(A, B, C)
	return C
}
