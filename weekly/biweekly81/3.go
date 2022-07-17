/**

 */

package weekly

import (
	"fmt"
	"sort"
)

func maximumXOR(nums []int) int {
	sort.Ints(nums)
	var answer = 0
	for i := 0; i < len(nums); i++ {
		answer = answer | nums[i]
	}
	return answer
	for i := 1; i < len(nums); i++ {
		answer = nums[i] ^ answer
		//fmt.Printf("%06b\n", nums[i])
	}
	c := fmt.Sprintf("%b", nums[len(nums)-1])
	hope := 1<<(len(c)) - 1
	if answer == hope {
		return answer
	}
	num := answer ^ hope
	// 增加一个数字，只能将1变成0
	//fmt.Printf("%06b\n", 1<<(len(c))-1)
	fmt.Printf("--%06b\n", answer^hope)
	var x, n = nums[0] - nums[0]&(nums[0]^num), 0
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i]&(nums[i]^num)
		if tmp > x {
			n = i
			x = tmp
		}
	}
	fmt.Println(n, nums[n], x, nums[n]&(nums[n]^num))
	fmt.Printf("--%b--\n--%b--\n", nums[n], nums[n]^num)
	fmt.Printf("--%b--\n--%b--", 805, 813)

	answer = nums[n]&(nums[n]^num) ^ answer ^ nums[n]
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
