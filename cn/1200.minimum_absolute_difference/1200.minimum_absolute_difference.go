//给你个整数数组 arr，其中每个元素都 不相同。
//
// 请你找到所有具有最小绝对差的元素对，并且按升序的顺序返回。
//
//
//
// 示例 1：
//
// 输入：arr = [4,2,1,3]
//输出：[[1,2],[2,3],[3,4]]
//
//
// 示例 2：
//
// 输入：arr = [1,3,6,10,15]
//输出：[[1,3]]
//
//
// 示例 3：
//
// 输入：arr = [3,8,-10,23,19,-4,-14,27]
//输出：[[-14,-10],[19,23],[23,27]]
//
//
//
//
// 提示：
//
//
// 2 <= arr.length <= 10^5
// -10^6 <= arr[i] <= 10^6
//
// 👍 109 👎 0

package cn

import (
	"math"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minimumAbsDifference(arr []int) (ans [][]int) {
	sort.Ints(arr)
	min := math.MaxInt
	for i := 1; i < len(arr); i++ { // 只有两个相邻的元素相减的绝对差最小
		a, b := arr[i], arr[i-1]
		v := a - b
		if v < min {
			ans = [][]int{{b, a}}
			min = v
		} else if min == v {
			ans = append(ans, []int{b, a})
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
