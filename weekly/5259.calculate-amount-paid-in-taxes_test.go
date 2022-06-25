package weekly

import "testing"

func Test_calculateTax(t *testing.T) {
	type args struct {
		brackets [][]int
		income   int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"", args{[][]int{{3, 50}, {7, 10}, {12, 25}}, 10}, 2.65},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTax(tt.args.brackets, tt.args.income); got != tt.want {
				t.Errorf("calculateTax() = %v, want %v", got, tt.want)
			}
		})
	}
}