package heap

import (
	"fmt"
	"testing"
)

func Test_newHeap(t *testing.T) {

}

func Test_newHeap1(t *testing.T) {
	arr := make([]int, 0)
	//for i := 0; i < 10; i++ {
	//	n := rand.Intn(10)
	//	arr = append(arr, n)
	//}
	/*for i := 10; i >= 0; i-- {
		arr = append(arr, i)
	}*/
	for i := 1; i < 11; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	h := NewHeap(arr, false)
	fmt.Println(h.getData())
}
func Test_newHeap2(t *testing.T) {
	arr := make([]int, 0)
	//for i := 0; i < 10; i++ {
	//	n := rand.Intn(10)
	//	arr = append(arr, n)
	//}
	/*for i := 10; i >= 0; i-- {
		arr = append(arr, i)
	}*/
	h := NewHeap(arr, false)
	a := []int{3, 8, 2, 5, 7, 9, 11}
	for i := 0; i < len(a); i++ {
		h.push(a[i])
	}
	var ans = make([]int, 0)
	for !h.isEmpty() {
		ans = append(ans, h.pop())
	}
	fmt.Println(ans)
}
