package cn

import "testing"

func Test_distributeCookies(t *testing.T) {
	type args struct {
		cookies []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{8, 15, 10, 20, 8}, 2}, 31},
		{"", args{[]int{6, 1, 3, 2, 2, 4, 1, 2}, 3}, 7},
		{"", args{[]int{76265, 7826, 16834, 63341, 68901, 58882, 50651, 75609}, 8}, 76265},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCookies(tt.args.cookies, tt.args.k); got != tt.want {
				t.Errorf("distributeCookies() = %v, want %v", got, tt.want)
			}
		})
	}
}
