package weekly

import "testing"

func Test_fillCups(t *testing.T) {
	type args struct {
		amount []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 4, 2}}, 4},
		{"", args{[]int{5, 4, 4}}, 7},
		{"", args{[]int{5, 0, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillCups(tt.args.amount); got != tt.want {
				t.Errorf("fillCups() = %v, want %v", got, tt.want)
			}
		})
	}
}
