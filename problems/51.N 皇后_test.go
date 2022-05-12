package problems

import (
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

		{"", args{1}, [][]string{{"Q"}}},
		{"", args{4}, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		{"8", args{8}, [][]string{{"aaa"}, {"aaaa"}}},
		{"12", args{12}, [][]string{{"aaa"}, {"aaaa"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildTree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{"4", args{4}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTree(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTree(t *testing.T) {
	n := 4
	tree := BuildTree(n)
	Lookup(tree)
}
