package problems

import "testing"

func Test_countElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{11, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 2, 15}}, 18},
		{"", args{[]int{11, 7, 2, 15}}, 2},
		{"", args{[]int{11}}, 0},
		{"", args{[]int{-3, 3, 3, 90}}, 2},
		{"", args{[]int{1, 1, 1, 1, 3, 3, 90}}, 2},
		{"", args{[]int{1, 1, 1, 1, 3, 3, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90}}, 2},
		{"", args{[]int{1, 1, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.nums); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
