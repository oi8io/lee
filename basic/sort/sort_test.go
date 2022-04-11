package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"偶数个",
			args{[]int{1, 3, 199, 12, 145, 45, 88, 99}}, []int{1, 3, 12, 45, 88, 99, 145, 199},
		},
		{
			"基数个",
			args{[]int{1, 3, 199, 12, 79, 145, 45, 88, 99}}, []int{1, 3, 12, 45, 79, 88, 99, 145, 199},
		},
		{
			"单个数字",
			args{[]int{99}}, []int{99},
		},
		{
			"两个数字",
			args{[]int{88, 99}}, []int{88, 99},
		},
		{
			"两个倒序数字",
			args{[]int{99, 88}}, []int{88, 99},
		},
		{
			"三个乱序数字",
			args{[]int{77, 99, 88}}, []int{77, 88, 99},
		},
		{
			"三个乱序数字",
			args{[]int{77, 99, 88, 199}}, []int{77, 88, 99, 199},
		},
		{
			"四个乱序数字",
			args{[]int{666, 77, 99, 88}}, []int{77, 88, 99, 666},
		},
		{
			"五个乱序数字",
			args{[]int{79, 145, 45, 88, 99}}, []int{45, 79, 88, 99, 145},
		},
		{
			"五个乱序数字",
			args{getInputs()}, getOutputs(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func GetMiddle() {

}

/**
1. 当i < j 则 i 前进
2. 当j
*/
func QuickSort1(in []int) (out []int) {
	if len(in) <= 1 {
		return in
	}
	// left exchange last
	if in[0] > in[len(in)-1] {
		in[0], in[len(in)-1] = Swap(in[0], in[len(in)-1])
	}
	//left exchange center
	if in[0] > in[len(in)/2] {
		in[0], in[len(in)/2] = Swap(in[0], in[len(in)/2])
	}
	//center exchange  last
	if in[len(in)/2] > in[len(in)-1] {
		in[len(in)/2], in[len(in)-1] = Swap(in[len(in)/2], in[len(in)-1])
	}

	var i, j = 0, len(in) - 1
	for {
		if i == j {
			break
		}
		if in[i] > in[j] {
			in[i], in[j] = Swap(in[i], in[j])
			j--
		} else {
			i++
		}
	}
	fmt.Println(in)
	return nil
}

func TestQuickSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []int
	}{
		// TODO: Add test cases.

		{
			"偶数个",
			args{[]int{1, 3, 199, 12, 145, 45, 88, 99}}, []int{1, 3, 12, 45, 88, 99, 145, 199},
		}, {
			"偶数个",
			args{[]int{1, 12, 3}}, []int{1, 3, 12},
		},
		{
			"基数个",
			args{[]int{1, 3, 199, 12, 79, 145, 45, 88, 99}}, []int{1, 3, 12, 45, 79, 88, 99, 145, 199},
		},
		{
			"基数个",
			args{[]int{1, 3, 99, 12, 79, 145, 45, 88, 199}}, []int{1, 3, 12, 45, 79, 88, 99, 145, 199},
		},
		{
			"单个数字",
			args{[]int{99}}, []int{99},
		},
		{
			"两个数字",
			args{[]int{88, 99}}, []int{88, 99},
		},
		{
			"两个倒序数字",
			args{[]int{99, 88}}, []int{88, 99},
		},
		{
			"三个乱序数字",
			args{[]int{77, 99, 88}}, []int{77, 88, 99},
		},
		{
			"三个乱序数字",
			args{[]int{5, 1, 1, 2, 0, 0}}, []int{5, 1, 1, 2, 0, 0},
		},
		{
			"三个乱序数字",
			args{[]int{88, 77, 99}}, []int{77, 88, 99},
		},
		{
			"四个乱序数字",
			args{[]int{666, 77, 99, 88}}, []int{77, 88, 99, 666},
		},
		{
			"五个乱序数字",
			args{getInputs()}, getOutputs(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := QuickSort(tt.args.in); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("QuickSort() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []int
	}{
		{
			"偶数个",
			args{[]int{1, 3, 199, 12, 145, 45, 88, 99}}, []int{1, 3, 12, 45, 88, 99, 145, 199},
		},
		{
			"基数个",
			args{[]int{1, 3, 199, 12, 79, 145, 45, 88, 99}}, []int{1, 3, 12, 45, 79, 88, 99, 145, 199},
		},
		{
			"单个数字",
			args{[]int{99}}, []int{99},
		},
		{
			"两个数字",
			args{[]int{88, 99}}, []int{88, 99},
		},
		{
			"两个倒序数字",
			args{[]int{99, 88}}, []int{88, 99},
		},
		{
			"三个乱序数字",
			args{[]int{77, 99, 88}}, []int{77, 88, 99},
		},
		{
			"四个乱序数字",
			args{[]int{666, 77, 99, 88}}, []int{77, 88, 99, 666},
		},
		{
			"五个乱序数字",
			args{[]int{79, 145, 45, 88, 99}}, []int{45, 79, 88, 99, 145},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := MergeSort(tt.args.in); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("MergeSort() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestSwapArr1(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"", args{1, 8}, 8, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Swap(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SwapArr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SwapArr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
