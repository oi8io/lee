package weekly

import "testing"

func Test_maximumXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{3, 2, 4, 6}}, 7},
		{"", args{[]int{1, 2, 3, 9, 2}}, 11},
		{"", args{[]int{772, 45, 1, 297, 549, 549, 301, 805, 297, 261, 36, 772, 37, 548, 300, 545, 773, 549, 268, 32, 41, 521, 44, 516, 256, 45, 549}}, 813},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumXOR(tt.args.nums); got != tt.want {
				t.Errorf("maximumXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
