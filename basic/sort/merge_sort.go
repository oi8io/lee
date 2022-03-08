package sort

// MergeSort 归并排序，核心采用线性合并两个已排序序列
func MergeSort(in []int) (out []int) {
	//
	if len(in) < 2 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = Swap(in[0], in[1])
		}
		return in
	}
	// 拆分成左右两部分
	ret1 := MergeSort(in[:len(in)/2])
	ret2 := MergeSort(in[len(in)/2:])
	out = Merge(ret1, ret2)
	//fmt.Println(in, out, len(in)/2, in[:len(in)/2], in[len(in)/2:], ret1, ret2)
	return out
}

func Merge(ret1 []int, ret2 []int) (out []int) {
	j, i, imax, jmax := 0, 0, len(ret1), len(ret2)
	// 循环，避免越界
	for i <= imax && j <= jmax {
		if i == imax { // 左边的数组遍历完成，把所有的右边剩余部分追加到数组
			out = append(out, ret2[j:]...)
			break
		}
		if j == jmax { // 同上
			out = append(out, ret1[i:]...)
			break
		}
		// 将较小的数字追加到数组
		if ret1[i] < ret2[j] {
			out = append(out, ret1[i])
			i++
		} else {
			out = append(out, ret2[j])
			j++
		}
	}
	return out
}

func Swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = b ^ a
	return a, b
}
