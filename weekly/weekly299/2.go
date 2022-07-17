package weekly

func countHousePlacements(n int) int {
	var dp = make([][]int, n)
	num := 1000000000 + 7
	dp[0] = make([]int, 2)
	dp[0][0] = 1
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 2)
		// 当前放置+当前不放置的方法总和
		dp[i][0] = (dp[i-1][1] + dp[i-1][0]) % num
		dp[i][1] = dp[i-1][0]
	}
	answer := (dp[n-1][0] + dp[n-1][1]) % num
	return answer * answer % num
}
