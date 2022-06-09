//二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。每个 LED 代表一个 0 或 1，最低位在右侧。
//
//
//
// 例如，下面的二进制手表读取 "3:25" 。
//
//
//
//
// （图源：WikiMedia - Binary clock samui moon.jpg ，许可协议：Attribution-ShareAlike 3.0
//Unported (CC BY-SA 3.0) ）
//
// 给你一个整数 turnedOn ，表示当前亮着的 LED 的数量，返回二进制手表可以表示的所有可能时间。你可以 按任意顺序 返回答案。
//
// 小时不会以零开头：
//
//
// 例如，"01:00" 是无效的时间，正确的写法应该是 "1:00" 。
//
//
// 分钟必须由两位数组成，可能会以零开头：
//
//
// 例如，"10:2" 是无效的时间，正确的写法应该是 "10:02" 。
//
//
//
//
// 示例 1：
//
//
//输入：turnedOn = 1
//输出：["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
//
//
// 示例 2：
//
//
//输入：turnedOn = 9
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= turnedOn <= 10
//
// 👍 364 👎 0

package cn

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func readBinaryWatch(turnedOn int) []string {
	var ans []string
	if turnedOn > 8 {
		return ans
	}
	if turnedOn == 0 {
		return []string{"0:00"}
	}
	backtrack(0, 0, turnedOn, &ans)
	return ans
}

/*// 伪码
var res [][]int
func backTrack(路径，选择列表){
	if 满足结束条件{ // 终止条件
		res = append(res, 路径)  // 存放结果
		return
	}
	for _,选择 := range 选择列表{ // 选择：本层集合中元素（树中节点孩子的数量就是集合的大小）
		做选择  // 处理节点
		backTrack(路径，选择列表)  // 递归
		撤销选择  // 回溯，撤销处理结果
	}
}
*/
func backtrack(p, v, cnt int, ans *[]string) {
	if cnt == 0 {
		//fmt.Printf("%02d result : %010b\n", p, v)
		*ans = append(*ans, GetBinTimeString(v))
		return
	}

	i := p
	for ; i < 10; i++ {
		if v&(i<<i) != 0 {
			continue
		}
		nv := (1 << i) | v
		h, m := getBinTime(nv)
		if h > 11 || m > 59 {
			continue
		} else {
			count := cnt
			np := i
			np++
			count--
			backtrack(np, nv, count, ans)
		}
	}
	return
}
func GetBinTimeString(val int) string {
	h, m := getBinTime(val)
	return fmt.Sprintf("%d:%02d", h, m)
}
func getBinTime(val int) (h int, m int) {
	return val >> 6, val & (1<<6 - 1)
}

//leetcode submit region end(Prohibit modification and deletion)
