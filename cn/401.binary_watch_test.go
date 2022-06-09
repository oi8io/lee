package cn

import (
	"reflect"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	type args struct {
		turnedOn int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{0}, []string{}},
		{"", args{1}, []string{}},
		{"", args{2}, []string{}},
		{"", args{3}, []string{}},
		{"", args{4}, []string{}},
		{"", args{5}, []string{}},
		{"", args{6}, []string{}},
		{"", args{7}, []string{}},
		{"", args{8}, []string{}},
		{"", args{9}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatch(tt.args.turnedOn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBinaryWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBinTime(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name  string
		args  args
		wantH int
		wantM int
	}{
		// TODO: Add test cases.
		{"", args{val: 518}, 1, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotH, gotM := getBinTime(tt.args.val)
			if gotH != tt.wantH {
				t.Errorf("getBinTime() gotH = %v, want %v", gotH, tt.wantH)
			}
			if gotM != tt.wantM {
				t.Errorf("getBinTime() gotM = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
