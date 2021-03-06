package cn

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"", args{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4}, []string{"the", "is", "sunny", "day"}},
		{"", args{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2}, []string{"i", "love"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
