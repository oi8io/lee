package sort

// InsertSort 插入排序，一开始手上的牌是空的，每次拉一张牌上来之后，从右往左比较大小，如果数字比它大，则向后移动，直到找到比其小的数字，将该位置
func InsertSort(nums []int) []int {
	var out []int
	for i := 0; i < len(nums); i++ {
		var pos = len(out)
		for j := 0; j < len(out); j++ {
			if nums[i] <= out[j] {
				pos = j
				//fmt.Println(pos, out[:pos], out[pos], nums[i], out[pos:])
				break
			}
		}
		var tmp = make([]int, 0)
		tmp = append(tmp, out[:pos]...)
		tmp = append(tmp, nums[i])
		out = append(tmp, out[pos:]...)
	}
	return out
}

func ShellSort(nums []int) []int {
	return nil
}

func BubbleSort(nums []int) []int {
	var out []int
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < i; j++ {

		}
	}
	return out
}
func HeapSort(nums []int) []int {
	var out []int
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < i; j++ {

		}
	}
	return out
}

func BucketSort(nums []int) []int {
	return nil

}
func QuickSort(nums []int) []int {
	QSort(nums, 0, len(nums)-1)
	return nums
}

func GetPivot(nums []int, left, right int) int {
	mid := (left + right) / 2
	if nums[left] > nums[right] {
		swapSlice(nums, left, right)
	}
	if nums[left] > nums[mid] {
		swapSlice(nums, left, mid)
	}
	// 保证最右边的是中间元素
	if nums[mid] < nums[right] {
		swapSlice(nums, mid, right)
	}
	return nums[right]
}

func swapSlice(nums []int, a, b int) {
	if a == b {
		return
	}
	nums[a], nums[b] = Swap(nums[a], nums[b])
}

func QSort(nums []int, left, right int) {
	if left == right {
		return
	}
	pivot := GetPivot(nums, left, right)
	i := left
	j := right - 1
	var split int
	for {
		if i > j {
			split = i
			break
		}

		if nums[i] < pivot {
			i++ // 如果左边一直小 就 一路走到黑
			continue
		}

		if nums[j] >= pivot {
			j-- // 如果右边一直大，也一直走到黑
			continue
		}

		//交换两个数字
		if i < j {
			swapSlice(nums, i, j)
		}
	}
	//fmt.Printf("%+v [%d:%d] <%d> <%d> ", nums, left, right, nums[right], nums[split])
	// 交换枢纽元与当前位置
	swapSlice(nums, right, split)
	//fmt.Printf(" %+v \n ", nums)
	//fmt.Println(nums, left, right, pivot, split, nums[split])

	//fmt.Println(nums[split], nums)
	if left < split-1 {
		QSort(nums, left, split-1)
	}
	if right > split+1 {
		QSort(nums, split+1, right)
	}
}

func QuickSort2(nums []int) []int {
	if len(nums) <= 10 {
		return InsertSort(nums)
	}
	var left, right []int
	var pivot = nums[0]
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if n > pivot {
			right = append(right, n)
		} else {
			left = append(left, n)
		}
	}
	right = QuickSort(right)
	left = QuickSort(left)
	out := append([]int{}, left...)
	out = append(out, pivot)
	out = append(out, right...)
	return out
}

// MergeSort 归并排序，核心采用线性合并两个已排序序列
func MergeSort(nums []int) (out []int) {
	//
	if len(nums) < 2 {
		return nums
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = Swap(nums[0], nums[1])
		}
		return nums
	}
	// 拆分成左右两部分
	ret1 := MergeSort(nums[:len(nums)/2])
	ret2 := MergeSort(nums[len(nums)/2:])
	out = Merge(ret1, ret2)
	//fmt.Println(nums, out, len(nums)/2, nums[:len(nums)/2], nums[len(nums)/2:], ret1, ret2)
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
