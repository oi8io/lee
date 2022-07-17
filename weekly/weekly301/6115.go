package weekly

import "fmt"

const N = 1000000007

func idealArrays(n int, maxValue int) int {
	var answer = 0
	for i := 1; i <= maxValue; i++ {
		x := maxValue / i
		//1开始
		ans := 1
		for j := 1; j < n; j++ {
			fmt.Println(j, ans)
			ans = (x * ans) % N
		}
		answer += ans / x
	}
	return answer
}
