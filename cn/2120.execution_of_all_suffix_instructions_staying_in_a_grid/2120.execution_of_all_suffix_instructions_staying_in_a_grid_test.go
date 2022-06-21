package cn

import (
	"reflect"
	"testing"
)

func Test_executeInstructions(t *testing.T) {
	type args struct {
		n        int
		startPos []int
		s        string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{3, []int{0, 1}, "RRDDLU"}, []int{1, 5, 4, 3, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeInstructions(tt.args.n, tt.args.startPos, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
