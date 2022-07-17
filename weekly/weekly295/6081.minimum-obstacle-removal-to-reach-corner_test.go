package weekly295

import "testing"

func Test_minimumObstacles(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}}, 2},
		{"", args{[][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumObstacles(tt.args.grid); got != tt.want {
				t.Errorf("minimumObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
