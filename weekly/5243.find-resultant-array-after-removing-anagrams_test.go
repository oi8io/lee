package weekly

import (
	"reflect"
	"testing"
)

func Test_removeAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"", args{[]string{"abba", "baba", "bbaa", "cd", "cd"}}, []string{"abba", "cd"}},
		{"", args{[]string{"a", "b", "c", "d", "e"}}, []string{"a", "b", "c", "d", "e"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeAnagrams(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
