package cn

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{[]int{3, 30, 34, 5, 9}}, "9534330"},
		{"", args{[]int{3, 30, 34, 5, 9}}, "9534330"},
		{"", args{[]int{3, 30}}, "330"},
		{"", args{[]int{3, 34}}, "343"},
		{"", args{[]int{10, 2}}, "210"},
		{"", args{[]int{10, 2, 9, 39, 17}}, "93921710"},
		{"", args{[]int{10, 17}}, "1710"},
		{"", args{[]int{10, 10, 17}}, "171010"},
		{"", args{[]int{8308, 8308, 830}}, "83088308830"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
