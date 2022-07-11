package cn

import (
	"fmt"
	"testing"
)

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"}, true},
		{"", args{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"}, true},
		{"", args{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"}, false},
		{"", args{[][]byte{{'A'}}, "B"}, false},
		{"", args{[][]byte{{'A'}}, "A"}, true},
		{"", args{[][]byte{{'a', 'a', 'b', 'a', 'a', 'b'}, {'a', 'a', 'b', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a', 'b', 'a'}, {'b', 'a', 'b', 'b', 'a', 'b'}, {'a', 'b', 'b', 'a', 'b', 'a'}, {'b', 'a', 'a', 'a', 'a', 'b'}}, "bbbaabbbbbab"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLocal(t *testing.T) {
	local := isLocal(2, 1, 2)
	fmt.Println(local)
}
