package cn

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"", args{4}, [][]string{{""}, {""}}},
		{"", args{8}, [][]string{{""}, {""}}},
		{"", args{10}, [][]string{{""}, {""}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	var x = make([]uint64, 16)
	for i := 0; i < 1024; i++ {
		//fmt.Println(i % 64)
		var n uint64 = 1 << (i % 64)
		x[i/64] = x[i/64] | n
	}
	for i := 0; i < len(x); i++ {
		fmt.Printf("%064b\n", x[i])
	}
}

func Test_printQueenBin(t *testing.T) {
	printQueenBin([]uint64{16612121211212}, 8)

}

func Test_checkBoard(t *testing.T) {
	type args struct {
		board []uint64
		n     int
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{[]uint64{1 << 1}, 8, 2, 1}, false},
		{"", args{[]uint64{1 << 0}, 8, 1, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkBoard(tt.args.board, tt.args.n, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("checkBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
