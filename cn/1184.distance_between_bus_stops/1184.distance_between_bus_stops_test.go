package cn

import "testing"

func Test_distanceBetweenBusStops(t *testing.T) {
	type args struct {
		distance    []int
		start       int
		destination int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 3, 4}, 1, 3}, 5},
		{"", args{[]int{1, 2, 3, 4}, 0, 1}, 1},
		{"", args{[]int{7, 10, 1, 12, 11, 14, 5, 0}, 7, 2}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceBetweenBusStops(tt.args.distance, tt.args.start, tt.args.destination); got != tt.want {
				t.Errorf("distanceBetweenBusStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
