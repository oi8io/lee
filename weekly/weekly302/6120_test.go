package weekly

import (
	"reflect"
	"testing"
)

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 3, 2, 1, 3, 2, 2}}, []int{3, 1}},
		{"", args{[]int{1, 1}}, []int{1, 0}},
		{"", args{[]int{0}}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPairs(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
