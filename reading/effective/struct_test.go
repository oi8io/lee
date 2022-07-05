package effective

import (
	"fmt"
	"testing"
)

func TestTester(t *testing.T) {
	ts := tester{s: "laiboamei"}
	tester.Print(ts, "lai le a ")
	(*tester).Print1(&ts, "lai le a ")
}

func TestSlice(t *testing.T) {

	arr := make([]int, 8)
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println(len(arr), cap(arr), arr)
	fmt.Println(len(arr), cap(arr))
	arr = make([]int, 16, 16)
	fmt.Println(arr[3])
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println(len(arr), cap(arr), arr)
	fmt.Println(len(arr), cap(arr))
	arr = make([]int, 0, 16) //
	arr = append(arr, 1, 2, 3, 4, 5)
	arr = append(arr, 1, 2, 3, 4, 5)
	arr = append(arr, 1, 2, 3, 4, 5)
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println(len(arr), cap(arr), arr)
	fmt.Println(len(arr), cap(arr))
}
