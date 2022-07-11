package weekly

import "testing"

func Test_validSubarraySize(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 3, 4, 3, 1}, 6}, 3},
		{"", args{[]int{6, 5, 6, 5, 8}, 7}, 1},
		{"", args{[]int{818, 232, 595, 418, 608, 229, 37, 330, 876, 774, 931, 939, 479, 884, 354, 328}, 3790}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validSubarraySize(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("validSubarraySize() = %v, want %v", got, tt.want)
			}
		})
	}
}
