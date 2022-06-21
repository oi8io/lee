/**
6096. 咒语和药水的成功对数 显示英文描述
通过的用户数388
尝试过的用户数537
用户总通过次数388
用户总提交次数620
题目难度Medium
给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[j] 表示第 j 瓶药水的能量强度。

同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合。

请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目。



示例 1：

输入：spells = [5,1,3], potions = [1,2,3,4,5], success = 7
输出：[4,0,3]
解释：
- 第 0 个咒语：5 * [1,2,3,4,5] = [5,10,15,20,25] 。总共 4 个成功组合。
- 第 1 个咒语：1 * [1,2,3,4,5] = [1,2,3,4,5] 。总共 0 个成功组合。
- 第 2 个咒语：3 * [1,2,3,4,5] = [3,6,9,12,15] 。总共 3 个成功组合。
所以返回 [4,0,3] 。
示例 2：

输入：spells = [3,1,2], potions = [8,5,8], success = 16
输出：[2,0,2]
解释：
- 第 0 个咒语：3 * [8,5,8] = [24,15,24] 。总共 2 个成功组合。
- 第 1 个咒语：1 * [8,5,8] = [8,5,8] 。总共 0 个成功组合。
- 第 2 个咒语：2 * [8,5,8] = [16,10,16] 。总共 2 个成功组合。
所以返回 [2,0,2] 。


提示：

n == spells.length
m == potions.length
1 <= n, m <= 105
1 <= spells[i], potions[i] <= 105
1 <= success <= 1010
*/

package weekly

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	var ans []int
	for i := 0; i < len(spells); i++ {
		i2 := pairs(potions, int64(spells[i]), success)
		ans = append(ans, i2)
	}
	return ans
}
func pairs(potions []int, cur int64, success int64) int {
	var start, end = 0, len(potions)
	if int64(potions[0])*cur >= success {
		return len(potions)
	}
	for {
		if start >= end-1 {
			break
		}
		mid := (start + end) / 2
		if int64(potions[mid])*cur < success {
			if start != end {
				start = mid
			} else {
				start++
			}
		} else {
			end = mid
		}
	}
	return len(potions) - (start + 1)
}
