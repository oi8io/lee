package weekly304

import "testing"

func Test_maximumGroups(t *testing.T) {
	type args struct {
		grades []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{8, 8}}, 1},
		{"", args{[]int{10, 6, 12, 7, 3, 5}}, 3},
		{"", args{[]int{8, 4, 4, 4, 6, 6, 6, 7, 7, 7}}, 3},
		{"", args{[]int{2, 31, 41, 31, 36, 46, 33, 45, 47, 8, 31, 6, 12, 12, 15, 35}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGroups(tt.args.grades); got != tt.want {
				t.Errorf("maximumGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
