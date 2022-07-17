package weekly

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"", args{3, [][]int{{0, 1}, {1, 2}, {2, 0}}}, 2},
		{"", args{6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}}, 44},
		{"", args{7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}}, 14},
		{"", args{3, [][]int{{0, 1}, {0, 2}, {1, 2}}}, 0},
		{"", args{4, nil}, 999000},
		{"", args{10, nil}, 999000},
		{"", args{1000, nil}, 999000},
		{"", args{100000, nil}, 4999950000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("validPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
