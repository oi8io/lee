package basic

import (
	"reflect"
	"testing"
)

func TestGetMinDistance(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name   string
		args   args
		wantP1 Point
		wantP2 Point
	}{
		{"", args{[]Point{{1, 2}, {3, 4}, {5, 6}}}, Point{1, 2}, Point{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1, gotP2 := GetMinDistance(tt.args.points)
			if !reflect.DeepEqual(gotP1, tt.wantP1) {
				t.Errorf("GetMinDistance() gotP1 = %v, want %v", gotP1, tt.wantP1)
			}
			if !reflect.DeepEqual(gotP2, tt.wantP2) {
				t.Errorf("GetMinDistance() gotP2 = %v, want %v", gotP2, tt.wantP2)
			}
		})
	}
}
