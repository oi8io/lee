package weekly294

import "testing"

func Test_maximumBags(t *testing.T) {
	type args struct {
		capacity        []int
		rocks           []int
		additionalRocks int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		//输入：capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2
		{"", args{[]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2}, 3},
		{"", args{[]int{10, 2, 2}, []int{2, 2, 0}, 100}, 3},
		{"", args{[]int{10, 1000000000}, []int{10, 0}, 1000000000}, 2},
		//输入：capacity = [10,2,2], rocks = [2,2,0], additionalRocks = 100

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBags(tt.args.capacity, tt.args.rocks, tt.args.additionalRocks); got != tt.want {
				t.Errorf("maximumBags() = %v, want %v", got, tt.want)
			}
		})
	}
}
