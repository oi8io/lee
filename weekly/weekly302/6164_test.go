package weekly

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums       []int
		numsDivide []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 3, 2, 4, 3}, []int{9, 6, 9, 3, 15}}, 2},
		{"", args{[]int{4, 3, 6}, []int{8, 2, 6, 10}}, -1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.numsDivide); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
