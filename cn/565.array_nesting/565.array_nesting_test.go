package cn

import "testing"

func Test_arrayNesting(t *testing.T) {
	type args struct {
		nums []int
	}
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}
	arr[9] = 0
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{5, 4, 0, 3, 1, 6, 2}}, 4},
		{"", args{arr}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayNesting(tt.args.nums); got != tt.want {
				t.Errorf("arrayNesting() = %v, want %v", got, tt.want)
			}
		})
	}
}
