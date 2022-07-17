package weekly294

import "testing"

func Test_minimumLines(t *testing.T) {
	type args struct {
		stockPrices [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		//输入：stockPrices = [[3,4],[1,2],[7,8],[2,3]]
		{"", args{[][]int{{1, 1}}}, 1},
		{"", args{[][]int{{3, 4}}}, 1},
		{"", args{[][]int{{3, 4}, {1, 2}, {7, 8}, {2, 3}}}, 1},
		{"", args{[][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}}}, 3},
		{"", args{[][]int{{57, 81}, {87, 91}, {64, 49}, {88, 49}, {45, 3}, {11, 63}, {61, 57}, {53, 63}, {94, 50}, {59, 9}, {44, 97}, {17, 29}, {19, 65}, {46, 77}, {82, 61}, {48, 86}, {73, 20}, {69, 29}, {49, 94}, {93, 56}, {98, 37}, {66, 91}, {90, 51}, {18, 2}, {76, 91}, {5, 6}, {33, 9}, {79, 23}, {52, 20}, {67, 13}, {85, 7}, {42, 71}, {30, 32}, {34, 25}, {29, 73}, {50, 27}, {80, 69}, {1, 59}, {31, 68}, {39, 8}, {99, 69}}}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLines(tt.args.stockPrices); got != tt.want {
				t.Errorf("minimumLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
