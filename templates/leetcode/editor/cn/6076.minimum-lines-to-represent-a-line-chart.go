/**
给你一个二维整数数组 stockPrices ，其中 stockPrices[i] = [dayi, pricei] 表示股票在 dayi 的价格为 pricei 。折线图 是一个二维平面上的若干个点组成的图，横坐标表示日期，纵坐标表示价格，折线图由相邻的点连接而成。比方说下图是一个例子：


请你返回要表示一个折线图所需要的 最少线段数 。



示例 1：



输入：stockPrices = [[1,7],[2,6],[3,5],[4,4],[5,4],[6,3],[7,2],[8,1]]
输出：3
解释：
上图为输入对应的图，横坐标表示日期，纵坐标表示价格。
以下 3 个线段可以表示折线图：
- 线段 1 （红色）从 (1,7) 到 (4,4) ，经过 (1,7) ，(2,6) ，(3,5) 和 (4,4) 。
- 线段 2 （蓝色）从 (4,4) 到 (5,4) 。
- 线段 3 （绿色）从 (5,4) 到 (8,1) ，经过 (5,4) ，(6,3) ，(7,2) 和 (8,1) 。
可以证明，无法用少于 3 条线段表示这个折线图。
示例 2：



输入：stockPrices = [[3,4],[1,2],[7,8],[2,3]]
输出：1
解释：
如上图所示，折线图可以用一条线段表示。


提示：

1 <= stockPrices.length <= 105
stockPrices[i].length == 2
1 <= dayi, pricei <= 109
所有 dayi 互不相同 。
*/

package problems

import (
	"sort"
)

func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) <= 0 {
		return 0
	}
	var days = make([]int, 0)
	var prices = make(map[int]int)
	for i := 0; i < len(stockPrices); i++ {
		day := stockPrices[i][0]
		price := stockPrices[i][1]
		days = append(days, day)
		prices[day] = price
	}
	sort.Ints(days)
	getRate := func(x1, x2 int) (n float64, m float64) {
		n = float64(prices[x2]-prices[x1]) / float64(x2-x1) // 小数情况
		m = float64(prices[x1]) - n*float64(x1)
		return
	}
	var n, m float64
	var count = 1
	var started bool
	for i := 1; i < len(days); i++ {
		n1, m1 := getRate(days[i-1], days[i])
		//fmt.Printf("{day:'%d',value:%d},", days[i], prices[days[i]])
		if started {
			if n1 != n || m1 != m {
				count++
			}
		}
		n, m = n1, m1
		started = true

	}
	return count
}
