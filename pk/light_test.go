package pk

import (
	"reflect"
	"testing"
)

func TestTurnOff(t *testing.T) {
	type args struct {
		light []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"5", args{[]int{1, 2, 3, 4, 5, 6, 7}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TurnOff(tt.args.light); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TurnOff() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Log("221313")
}
