package weekly

import "testing"

func Test_checkXMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		//	{{2,0,0,1},{0,3,1,0},{0,5,2,0},{4,0,0,2}},{{5,7,0},{0,3,1},{0,5,0}}
		{"", args{[][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}}, true},
		{"", args{[][]int{{5, 7, 0}, {0, 3, 1}, {0, 5, 0}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkXMatrix(tt.args.grid); got != tt.want {
				t.Errorf("checkXMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
