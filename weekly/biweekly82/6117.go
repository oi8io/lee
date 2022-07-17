package biweekly82

import (
	"sort"
)

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	cnt := 0
	end := 0
	for b := 0; b < len(buses); b++ {
		end = findNumber(passengers, buses[b])
		if end-cnt < capacity {
			for i := end; i >= cnt; i++ {

			}
			cnt += end + 1
		} else {
			cnt += capacity
		}
	}

	return 0
}

func findNumber(arr []int, n int) int {
	if len(arr) == 0 {
		return 0
	}
	min, max, mid := 0, len(arr), 0
	for {
		mid = min + (max-min)>>1
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
	if arr[mid] <= n {
		return mid
	} else {
		return mid - 1
	}
}
