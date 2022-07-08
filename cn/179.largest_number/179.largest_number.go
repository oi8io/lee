//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//
//
// 示例 1：
//
//
//输入：nums = [10,2]
//输出："210"
//
// 示例 2：
//
//
//输入：nums = [3,30,34,5,9]
//输出："9534330"
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 10⁹
//
// 👍 965 👎 0

package cn

import (
	"sort"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	x := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		x = append(x, strconv.Itoa(nums[i]))
	}
	sort.Slice(x, func(i, j int) bool {
		if x[i]+x[j] > x[j]+x[i] {
			return true
		} else {
			return false
		}
	})
	str := strings.Join(x, "")
	i := 0
	for ; i < len(str)-1; i++ {
		if str[i] != '0' {
			break
		}
	}
	return str[i:]
}

//leetcode submit region end(Prohibit modification and deletion)
