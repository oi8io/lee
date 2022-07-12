package cn

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantIdx int
	}{
		// TODO: Add test cases.
		{"", args{[]int{4, 3}}, 0},
		{"", args{[]int{3, 4}}, 1},
		{"", args{[]int{3, 5, 4}}, 1},
		{"", args{[]int{3, 1, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIdx := findPeakElement(tt.args.nums); gotIdx != tt.wantIdx {
				t.Errorf("findPeakElement() = %v, want %v", gotIdx, tt.wantIdx)
			}
		})
	}
}
