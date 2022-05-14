package problems

import (
	"reflect"
	"testing"
)

func Test_getDistances(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 1, 1, 1, 1, 1, 1}}, []int64{21, 16, 13, 12, 13, 16, 21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistances(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
