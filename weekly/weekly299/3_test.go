package weekly

import "testing"

func Test_maximumsSplicedArray(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{60, 60, 60}, []int{10, 90, 10}}, 210},
		{"", args{[]int{20, 40, 20, 70, 30}, []int{50, 20, 50, 40, 20}}, 220},
		{"", args{[]int{7, 11, 13}, []int{1, 1, 1}}, 31},
		{"", args{[]int{1, 1, 1, 1, 1, 13}, []int{1, 19, 1, 1, 11, 1}}, 46},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumsSplicedArray(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maximumsSplicedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
