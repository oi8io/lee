package link

func insert(arr []int, i, x int) []int {
	if i > 0 && arr[i-1] == x {
		return arr
	}
	nArr := make([]int, len(arr)+1)
	copy(nArr, arr[:i])
	copy(nArr[i+1:], arr[i:])
	nArr[i] = x
	return nArr
}

func findNumber(arr []int, n int) int {
	if len(arr) == 0 {
		return 0
	}
	min, max, mid := 0, len(arr), 0
	for {
		mid = (min + max) / 2
		if min >= max-1 {
			break
		}
		if arr[mid] > n {
			max = mid
		} else {
			if min == mid {
				min++
			} else {
				min = mid
			}
		}
	}
	if arr[mid] > n {
		return mid
	} else {
		return mid + 1
	}
}
