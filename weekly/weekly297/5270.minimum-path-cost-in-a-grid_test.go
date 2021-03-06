package weekly297

import "testing"

//go:lint  [[5,1,2],[4,0,3]], moveCost = [[12,10,15],[20,23,8],[21,7,1],[8,1,13],[9,10,25],[5,3,2]]
func Test_minPathCost(t *testing.T) {
	type args struct {
		grid     [][]int
		moveCost [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[][]int{{5, 3}, {4, 0}, {2, 1}}, [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}}, 17},
		{"", args{[][]int{{5, 1, 2}, {4, 0, 3}}, [][]int{{12, 10, 15}, {20, 23, 8}, {21, 7, 1}, {8, 1, 13}, {9, 10, 25}, {5, 3, 2}}}, 6},
		//[[28,35,29,5,13,17,18,23,14],[30,15,12,27,2,26,25,19,7],[1,16,34,21,9,3,20,24,8],[4,33,22,11,31,0,6,10,32]]
		//[[87,46,11,33,55,26,26,56,23],[77,56,72,49,35,18,37,66,37],[54,40,62,1,64,49,95,81,77],[80,7,76,71,91,67,75,84,52],[65,11,13,15,9,34,10,98,20],[1,95,100,61,33,47,28,100,44],[39,56,94,7,64,91,66,34,70],[37,99,62,7,23,78,74,89,97],[84,41,63,42,84,15,46,31,11],[60,36,27,25,37,18,4,90,43],[35,83,90,37,91,27,61,99,53],[85,2,98,99,67,70,38,91,68],[66,46,7,99,26,81,95,51,51],[41,96,66,84,61,73,78,28,63],[38,34,49,55,35,29,93,5,28],[3,30,80,20,23,10,93,40,33],[8,86,47,15,45,84,47,19,58],[72,5,76,82,60,50,13,74,38],[4,8,25,38,29,4,60,81,21],[65,50,74,32,9,47,71,55,14],[90,30,92,51,45,51,4,85,22],[30,56,1,45,63,72,91,73,60],[51,61,53,49,44,99,11,5,3],[24,54,91,11,5,30,50,89,44],[87,97,46,92,93,41,64,73,15],[94,76,90,80,30,9,88,8,33],[50,34,4,63,49,90,46,55,16],[10,46,80,21,97,69,70,85,31],[10,66,74,43,65,45,85,34,91],[82,26,77,10,2,5,89,39,4],[80,70,89,73,54,61,100,89,23],[30,66,80,51,3,34,92,100,63],[74,15,4,33,37,3,87,76,92],[44,43,77,99,27,1,23,10,8],[8,74,17,35,31,84,97,98,34],[99,9,28,43,55,39,93,64,93]]
		{"", args{[][]int{{28, 35, 29, 5, 13, 17, 18, 23, 14}, {30, 15, 12, 27, 2, 26, 25, 19, 7}, {1, 16, 34, 21, 9, 3, 20, 24, 8}, {4, 33, 22, 11, 31, 0, 6, 10, 32}}, [][]int{{87, 46, 11, 33, 55, 26, 26, 56, 23}, {77, 56, 72, 49, 35, 18, 37, 66, 37}, {54, 40, 62, 1, 64, 49, 95, 81, 77}, {80, 7, 76, 71, 91, 67, 75, 84, 52}, {65, 11, 13, 15, 9, 34, 10, 98, 20}, {1, 95, 100, 61, 33, 47, 28, 100, 44}, {39, 56, 94, 7, 64, 91, 66, 34, 70}, {37, 99, 62, 7, 23, 78, 74, 89, 97}, {84, 41, 63, 42, 84, 15, 46, 31, 11}, {60, 36, 27, 25, 37, 18, 4, 90, 43}, {35, 83, 90, 37, 91, 27, 61, 99, 53}, {85, 2, 98, 99, 67, 70, 38, 91, 68}, {66, 46, 7, 99, 26, 81, 95, 51, 51}, {41, 96, 66, 84, 61, 73, 78, 28, 63}, {38, 34, 49, 55, 35, 29, 93, 5, 28}, {3, 30, 80, 20, 23, 10, 93, 40, 33}, {8, 86, 47, 15, 45, 84, 47, 19, 58}, {72, 5, 76, 82, 60, 50, 13, 74, 38}, {4, 8, 25, 38, 29, 4, 60, 81, 21}, {65, 50, 74, 32, 9, 47, 71, 55, 14}, {90, 30, 92, 51, 45, 51, 4, 85, 22}, {30, 56, 1, 45, 63, 72, 91, 73, 60}, {51, 61, 53, 49, 44, 99, 11, 5, 3}, {24, 54, 91, 11, 5, 30, 50, 89, 44}, {87, 97, 46, 92, 93, 41, 64, 73, 15}, {94, 76, 90, 80, 30, 9, 88, 8, 33}, {50, 34, 4, 63, 49, 90, 46, 55, 16}, {10, 46, 80, 21, 97, 69, 70, 85, 31}, {10, 66, 74, 43, 65, 45, 85, 34, 91}, {82, 26, 77, 10, 2, 5, 89, 39, 4}, {80, 70, 89, 73, 54, 61, 100, 89, 23}, {30, 66, 80, 51, 3, 34, 92, 100, 63}, {74, 15, 4, 33, 37, 3, 87, 76, 92}, {44, 43, 77, 99, 27, 1, 23, 10, 8}, {8, 74, 17, 35, 31, 84, 97, 98, 34}, {99, 9, 28, 43, 55, 39, 93, 64, 93}}}, 59},
		{"", args{[][]int{{37, 30, 26, 1, 0, 42, 33}, {28, 39, 45, 54, 55, 34, 3}, {50, 17, 6, 41, 4, 18, 15}, {22, 32, 16, 44, 31, 40, 29}, {24, 2, 51, 46, 35, 27, 38}, {20, 25, 11, 49, 12, 13, 10}, {47, 14, 53, 5, 36, 43, 21}, {19, 9, 48, 8, 7, 23, 52}}, [][]int{{42, 31, 20, 33, 54, 87, 76}, {51, 26, 21, 77, 91, 4, 100}, {45, 81, 4, 32, 54, 47, 47}, {41, 25, 68, 100, 1, 19, 76}, {47, 69, 93, 66, 5, 23, 93}, {100, 63, 55, 8, 3, 41, 88}, {46, 59, 42, 79, 51, 12, 21}, {41, 22, 18, 5, 59, 66, 69}, {12, 19, 53, 28, 50, 74, 31}, {16, 89, 46, 45, 95, 39, 22}, {37, 37, 3, 93, 45, 82, 72}, {31, 61, 84, 38, 75, 70, 2}, {87, 30, 91, 38, 20, 33, 43}, {38, 19, 12, 33, 52, 26, 6}, {23, 22, 26, 37, 26, 55, 8}, {77, 7, 20, 10, 10, 69, 82}, {23, 21, 16, 24, 99, 58, 52}, {61, 55, 81, 38, 18, 38, 20}, {56, 97, 20, 80, 56, 22, 63}, {12, 1, 100, 80, 85, 70, 28}, {50, 89, 27, 61, 37, 63, 33}, {72, 49, 85, 93, 53, 91, 84}, {64, 91, 78, 69, 82, 85, 71}, {8, 50, 81, 21, 25, 21, 75}, {69, 34, 8, 82, 98, 5, 74}, {86, 37, 43, 95, 41, 27, 99}, {33, 85, 72, 75, 13, 99, 18}, {76, 84, 94, 17, 9, 17, 56}, {56, 67, 16, 98, 56, 55, 27}, {96, 79, 23, 95, 64, 23, 58}, {81, 2, 73, 81, 6, 42, 20}, {38, 68, 98, 83, 83, 60, 19}, {42, 85, 84, 99, 56, 53, 25}, {29, 64, 6, 18, 66, 24, 57}, {4, 87, 6, 62, 49, 42, 81}, {38, 24, 83, 10, 86, 53, 34}, {48, 44, 2, 38, 85, 18, 34}, {2, 79, 59, 96, 80, 53, 77}, {17, 1, 40, 84, 53, 21, 85}, {72, 15, 55, 59, 77, 64, 19}, {51, 3, 38, 36, 29, 18, 80}, {36, 67, 6, 32, 48, 74, 17}, {25, 13, 17, 84, 2, 39, 44}, {8, 60, 73, 20, 29, 59, 44}, {98, 22, 52, 4, 27, 55, 49}, {90, 50, 75, 98, 17, 87, 4}, {17, 21, 50, 93, 72, 34, 64}, {48, 88, 27, 69, 30, 88, 78}, {42, 70, 10, 2, 83, 71, 40}, {6, 5, 43, 25, 1, 90, 41}, {9, 93, 98, 41, 23, 25, 35}, {66, 93, 37, 23, 17, 73, 4}, {97, 1, 5, 92, 17, 28, 22}, {67, 49, 78, 74, 80, 30, 23}, {91, 61, 92, 61, 95, 81, 18}, {71, 3, 86, 51, 43, 91, 58}}}, 176},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathCost(tt.args.grid, tt.args.moveCost); got != tt.want {
				t.Errorf("minPathCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
