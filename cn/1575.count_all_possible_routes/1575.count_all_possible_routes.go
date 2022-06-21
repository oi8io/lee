//给你一个 互不相同 的整数数组，其中 locations[i] 表示第 i 个城市的位置。同时给你 start，finish 和 fuel 分别表示出发城市
//、目的地城市和你初始拥有的汽油总量
//
// 每一步中，如果你在城市 i ，你可以选择任意一个城市 j ，满足 j != i 且 0 <= j < locations.length ，并移动到城市
//j 。从城市 i 移动到 j 消耗的汽油量为 |locations[i] - locations[j]|，|x| 表示 x 的绝对值。
//
// 请注意， fuel 任何时刻都 不能 为负，且你 可以 经过任意城市超过一次（包括 start 和 finish ）。
//
// 请你返回从 start 到 finish 所有可能路径的数目。
//
// 由于答案可能很大， 请将它对 10^9 + 7 取余后返回。
//
//
//
// 示例 1：
//
//
//输入：locations = [2,3,6,8,4], start = 1, finish = 3, fuel = 5
//输出：4
//解释：以下为所有可能路径，每一条都用了 5 单位的汽油：
//1 -> 3
//1 -> 2 -> 3
//1 -> 4 -> 3
//1 -> 4 -> 2 -> 3
//
//
// 示例 2：
//
//
//输入：locations = [4,3,1], start = 1, finish = 0, fuel = 6
//输出：5
//解释：以下为所有可能的路径：
//1 -> 0，使用汽油量为 fuel = 1
//1 -> 2 -> 0，使用汽油量为 fuel = 5
//1 -> 2 -> 1 -> 0，使用汽油量为 fuel = 5
//1 -> 0 -> 1 -> 0，使用汽油量为 fuel = 3
//1 -> 0 -> 1 -> 0 -> 1 -> 0，使用汽油量为 fuel = 5
//
//
// 示例 3：
//
//
//输入：locations = [5,2,1], start = 0, finish = 2, fuel = 3
//输出：0
//解释：没有办法只用 3 单位的汽油从 0 到达 2 。因为最短路径需要 4 单位的汽油。
//
//
//
// 提示：
//
//
// 2 <= locations.length <= 100
// 1 <= locations[i] <= 10⁹
// 所有 locations 中的整数 互不相同 。
// 0 <= start, finish < locations.length
// 1 <= fuel <= 200
//
// 👍 69 👎 0

package cn

import (
	"fmt"
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)
func countRoutes(locations []int, start int, finish int, fuel int) int {
	/**
	1. 确定状态，起始位置，结束位置，油耗量
	2. 转移方程 dp[j]=fuel-|dp[i]-dp[i]|
	3. 初始条件，起始位置，结束位置，油耗量 ,边界条件，|dp[i]-dp[i]|<fluel 终点是finish
	4. 计算顺序，从左向右 全排列，先记录每个点到终点的油耗
	*/
	/**
	【4，3，1】
	                  ┌───────┐
	                  │  end  │
	          ┌───────┼───────┼───────┐
	          │   0   │   1   │   2   │
	       ┌──┼───────┼───────┼───────┤
	       │  │       │       │       │
	       │0 │   0   │   1   │   3   │
	       │  │       │       │       │
	┌──────┼──┼───────┼───────┼───────┤
	│start │  │       │       │       │
	└──────┤1 │   1   │   0   │   2   │
	       │  │       │       │       │
	       ├──┼───────┼───────┼───────┤
	       │  │       │       │       │
	       │2 │   3   │   2   │   0   │
	       │  │       │       │       │
	       └──┴───────┴───────┴───────┘
	图表表示，纵坐标表示开始，横坐标表示end所需要的油耗，只能从一个城市到另外一个城市，否则无意义
	如：0出发到1 需要耗油1
	*/
	// 油耗
	var costMap = make(map[int]int)
	for i := 0; i < len(locations); i++ {
		if i != finish {
			costMap[i] = getFuelCost(locations, i, finish)
		}
	}
	var routes int
	/*
		countRoutes1(locations, start, finish, fuel, costMap, &routes)
	*/
	routes = countRoutes3(locations, start, finish, fuel)

	if start == finish {
		return routes + 1
	}
	return routes

}

// 动态规划 fixme  失败了 失败了啊
func countRoutes3(locations []int, start int, finish int, fuel int) int {
	n := len(locations)
	mod := int(math.Pow(10, 9)) + 7
	dp := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := dp[i]; !ok {
			dp[i] = make(map[int]int)
		}
	}
	for i := 0; i <= fuel; i++ {
		dp[finish][i] = 1 //不管剩余多少油，
	}
	for c := 0; c <= fuel; c++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if j != i { //尝试各种路径
					cost := getFuelCost(locations, start, j)
					if c >= cost {
						dp[i][c] += dp[j][c-cost]
						dp[i][c] %= mod
					}
				}
			}
		}
	}
	fmt.Println(dp[finish])
	return dp[start][fuel]
}

// 动态规划 记忆化搜索
func countRoutes2(locations []int, start int, finish int, fuel int, dp map[int]map[int]int, costMap map[int]int) int {
	if _, ok := dp[start]; !ok {
		dp[start] = make(map[int]int)
	}
	if v, ok := dp[start][fuel]; ok {
		return v
	}
	ans := 0
	// 无效情况
	if fuel < costMap[start] || fuel == 0 {
		dp[start][fuel] = 0
		return 0
	}
	for i := 0; i < len(locations); i++ {
		if i == start { //尝试各种路径
			continue
		}
		cost := getFuelCost(locations, start, i)
		if cost > fuel {
			continue
		}
		if i == finish { // 统一这里累加
			ans++
		}
		nfuel := fuel - cost
		if nfuel > 0 {
			ans += countRoutes2(locations, i, finish, nfuel, dp, costMap)
		}
	}
	ans = ans % (int(math.Pow(10, 9)) + 7)
	dp[start][fuel] = ans
	return ans
}

// 回溯法
func countRoutes1(locations []int, start int, finish int, fuel int, costMap map[int]int, routes *int) {
	if start == finish {
		*routes++
		if fuel == 0 {
			return
		}
	}
	if fuel < costMap[start] || fuel == 0 {
		return
	}
	for i := 0; i < len(locations); i++ {
		if i == start { //尝试各种路径
			continue
		}
		cost := getFuelCost(locations, start, i)
		if cost > fuel {
			continue
		}
		nfuel := fuel - cost
		if nfuel > 0 {
			countRoutes1(locations, i, finish, nfuel, costMap, routes)
		}
	}
}

func getFuelCost(locations []int, start, end int) int {
	if locations[start] > locations[end] {
		return locations[start] - locations[end]
	}
	return locations[end] - locations[start]
}

//leetcode submit region end(Prohibit modification and deletion)
