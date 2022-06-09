//给你一个下标从 0 开始、由 n 个整数组成的数组 arr 。
//
// arr 中两个元素的 间隔 定义为它们下标之间的 绝对差 。更正式地，arr[i] 和 arr[j] 之间的间隔是 |i - j| 。
//
// 返回一个长度为 n 的数组 intervals ，其中 intervals[i] 是 arr[i] 和 arr 中每个相同元素（与 arr[i] 的值相同
//）的 间隔之和 。
//
// 注意：|x| 是 x 的绝对值。
//
//
//
// 示例 1：
//
// 输入：arr = [2,1,3,1,2,3,3]
//输出：[4,2,7,2,4,4,5]
//解释：
//- 下标 0 ：另一个 2 在下标 4 ，|0 - 4| = 4
//- 下标 1 ：另一个 1 在下标 3 ，|1 - 3| = 2
//- 下标 2 ：另两个 3 在下标 5 和 6 ，|2 - 5| + |2 - 6| = 7
//- 下标 3 ：另一个 1 在下标 1 ，|3 - 1| = 2
//- 下标 4 ：另一个 2 在下标 0 ，|4 - 0| = 4
//- 下标 5 ：另两个 3 在下标 2 和 6 ，|5 - 2| + |5 - 6| = 4
//- 下标 6 ：另两个 3 在下标 2 和 5 ，|6 - 2| + |6 - 5| = 5
//
//
// 示例 2：
//
// 输入：arr = [10,5,10,10]
//输出：[5,0,3,4]
//解释：
//- 下标 0 ：另两个 10 在下标 2 和 3 ，|0 - 2| + |0 - 3| = 5
//- 下标 1 ：只有这一个 5 在数组中，所以到相同元素的间隔之和是 0
//- 下标 2 ：另两个 10 在下标 0 和 3 ，|2 - 0| + |2 - 3| = 3
//- 下标 3 ：另两个 10 在下标 0 和 2 ，|3 - 0| + |3 - 2| = 4
//
//
//
//
// 提示：
//
//
// n == arr.length
// 1 <= n <= 10⁵
// 1 <= arr[i] <= 10⁵
//
// 👍 38 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func getDistances(arr []int) []int64 {
	numMaps := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		numMaps[arr[i]] = append(numMaps[arr[i]], i)
	}
	var result = make([]int64, len(arr))
	for _, p := range numMaps {
		sum := int64(0)
		for _, i := range p {
			sum += int64(i - p[0]) // 该组第一个元素的间隔和
		}
		result[p[0]] = sum
		for i, n := 1, len(p); i < n; i++ {
			sum += int64(2*i-n) * int64(p[i]-p[i-1]) // 计算该组下一个元素的间隔和（考虑变化量）
			result[p[i]] = sum
		}
	}
	return result
}

func absMinus(a, b int64) int64 {
	if a > b {
		return a - b
	}
	return b - a
}

//leetcode submit region end(Prohibit modification and deletion)
