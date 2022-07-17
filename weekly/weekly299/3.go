package weekly

import "fmt"

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	n := len(nums1)
	sum1, sum2 := 0, 0
	dp1, dp2 := make([][]int, n), make([][]int, n)
	//1, 1, 1, 1, 1, 13  1
	//1, 19, 1, 1, 11, 1 99
	//0  ,18,  18,18,28,28-12  98+28-12
	//                  12
	for i := 0; i < n; i++ {
		dp1[i], dp2[i] = make([]int, 2), make([]int, 2)
		n1, n2 := nums1[i], nums2[i]
		sum1 += n1
		sum2 += n2
		if i == 0 {
			dp1[i][0] = 0
			dp1[i][1] = n1 - n2
			dp2[i][0] = 0
			dp2[i][1] = n2 - n1
		} else {
			dp1[i][0] = max(dp1[i-1][0], dp1[i-1][1])
			dp1[i][1] = max(dp1[i-1][1]+n1-n2, n1-n2)
			dp2[i][0] = max(dp2[i-1][0], dp2[i-1][1])
			dp2[i][1] = max(dp2[i-1][1]+n2-n1, n2-n1)
		}
		//不处理当前位
	}
	fmt.Println(dp2, dp1)
	fmt.Println(sum2, sum1)
	ans1 := max(dp2[n-1][0], dp2[n-1][1]) + sum1
	ans2 := max(dp1[n-1][0], dp1[n-1][1]) + sum2
	return max(ans1, ans2)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
