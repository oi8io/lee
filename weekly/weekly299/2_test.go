package weekly

import "testing"

func Test_countHousePlacements(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{1}, 4},
		{"", args{2}, 9},
		{"", args{3}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHousePlacements(tt.args.n); got != tt.want {
				t.Errorf("countHousePlacements() = %v, want %v", got, tt.want)
			}
		})
	}
}
