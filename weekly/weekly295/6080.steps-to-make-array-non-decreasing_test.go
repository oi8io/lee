package weekly295

import "testing"

func Test_totalSteps(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}}, 3},
		{"", args{[]int{4, 5, 7, 7, 13}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalSteps(tt.args.nums); got != tt.want {
				t.Errorf("totalSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
