package stack

import (
	"reflect"
	"testing"
)

func TestMakeMonotonousStack(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1}}, [][]int{{-1, -1}}},
		{"", args{[]int{1, 2}}, [][]int{{-1, -1}, {0, -1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeMonotonousStack(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMonotonousStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
