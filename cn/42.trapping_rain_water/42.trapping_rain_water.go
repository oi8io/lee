//<p>给定&nbsp;<code>n</code> 个非负整数表示每个宽度为 <code>1</code> 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。</p>
//
//<p>&nbsp;</p>
//
//<p><strong>示例 1：</strong></p>
//
//<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png" style="height: 161px; width: 412px;" /></p>
//
//<pre>
//<strong>输入：</strong>height = [0,1,0,2,1,0,1,3,2,1,2,1]
//<strong>输出：</strong>6
//<strong>解释：</strong>上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//</pre>
//
//<p><strong>示例 2：</strong></p>
//
//<pre>
//<strong>输入：</strong>height = [4,2,0,3,2,5]
//<strong>输出：</strong>9
//</pre>
//
//<p>&nbsp;</p>
//
//<p><strong>提示：</strong></p>
//
//<ul>
//	<li><code>n == height.length</code></li>
//	<li><code>1 &lt;= n &lt;= 2 * 10<sup>4</sup></code></li>
//	<li><code>0 &lt;= height[i] &lt;= 10<sup>5</sup></code></li>
//</ul>
//<div><li>👍 3414</li><li>👎 0</li></div>

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	return trap3(height)
	return trap4(height)
	return trap2(height)
	return trap1(height)
}

// 双指针
func trap4(height []int) int {
	var lo, hi = 0, len(height) - 1
	var ans, ml, mr = 0, 0, 0

	for lo <= hi {
		l, r := height[lo], height[hi]
		if ml < l {
			ml = l
		} else {
			ans += ml - l
		}
		if mr < r {
			mr = r
		} else {
			ans += mr - r
		}
		if ml <= mr {
			lo++
		} else {
			hi--
		}
	}
	return ans
}

// 单调栈
func trap3(height []int) (ans int) {
	var stack = make([]int, 0)
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}

//dp 动态规划求每个位置左右最高
func trap2(height []int) int {
	n := len(height)
	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		left[i] = max(height[i-1], left[i-1])
		j := n - 1 - i
		right[j] = max(height[j+1], right[j+1])
	}
	ans := 0
	for i := n - 2; i >= 0; i-- {
		h := min(right[i], left[i])
		if height[i] < h {
			ans += h - height[i]
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trap1(height []int) int {
	n := len(height)
	h := 0
	for i := 1; i < n; i++ {
		if height[i] > height[h] {
			h = i
		}
	}

	ans, lh, rh := 0, 0, 0
	for i := 0; i < h; i++ {
		c := height[i]
		if c > lh {
			lh = c
		} else {
			ans += lh - c
		}
	}
	for i := n - 1; i > h; i-- {
		c := height[i]
		if c > rh {
			rh = c
		} else {
			ans += rh - c
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
