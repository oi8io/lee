package cn

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{"foofoofoofoofoofoo", []string{"foo"}}, []int{0, 3, 6, 9, 12, 15}},
		{"", args{"ofofofofofofofofofofofofofofo", []string{"ofo"}}, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26}},
		{"", args{"foofoofoofoofoofoofoo", []string{"foo", "foo", "foo"}}, []int{0, 3, 6, 9, 12}},
		{"", args{"foofoofoofoofoofoofoo", []string{"foo", "foo"}}, []int{0, 3, 6, 9, 12, 15}},
		{"", args{"barfoothefoobarman", []string{"foo", "bar"}}, []int{0, 9}},
		{"", args{"barfoobar", []string{"foo", "bar"}}, []int{0, 3}},
		{"", args{"barpbarxxxxx", []string{"bar", "arp"}}, []int{1}},
		{"", args{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}}, nil},
		{"", args{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}}, []int{6, 9, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
