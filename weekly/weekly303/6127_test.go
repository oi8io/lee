package weekly

import "testing"

func Test_countExcellentPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"5", args{[]int{1, 2, 3, 1}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countExcellentPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countExcellentPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
