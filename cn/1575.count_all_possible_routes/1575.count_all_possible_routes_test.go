package cn

import "testing"

func Test_countRoutes(t *testing.T) {
	type args struct {
		locations []int
		start     int
		finish    int
		fuel      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{2, 3, 6, 8, 4}, 1, 3, 5}, 4},
		{"", args{[]int{4, 3, 1}, 1, 0, 6}, 5},
		{"", args{[]int{2, 1, 5}, 0, 0, 3}, 2},
		{"", args{[]int{5, 2, 1}, 0, 2, 3}, 0},
		{"", args{[]int{1, 2, 3}, 0, 2, 10}, 182},
		{"", args{[]int{1, 2, 3}, 0, 2, 20}, 44286},
		{"", args{[]int{1, 2, 3}, 0, 2, 30}, 10761680},
		{"", args{[]int{1, 2, 3}, 0, 2, 40}, 615088286},
		{"", args{[]int{1, 2, 3}, 0, 2, 80}, 464953390},
		{"", args{[]int{1, 2, 3}, 0, 2, 100}, 32578211},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRoutes(tt.args.locations, tt.args.start, tt.args.finish, tt.args.fuel); got != tt.want {
				t.Errorf("countRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
