package cn

import (
	"fmt"
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	var ag1, ag2 []int
	for i := 1; i <= 64; i++ {
		if i == 3 {
			ag1 = append(ag1, 198)
			continue

		}
		ag1 = append(ag1, i)
	}
	for i := 1; i <= 127; i++ {
		ag2 = append(ag2, i)
	}
	fmt.Println(ag1)
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{[]int{1, 2, 3, 4, 32, 63, 64, 65, 99, 100005, 102121213121}}, want: 5},
		{name: "", args: args{[]int{999, 500, 1}}, want: 2},
		{name: "", args: args{[]int{1, 2, 0}}, want: 3},
		{name: "", args: args{ag1}, want: 3},
		{name: "", args: args{ag2}, want: 128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
