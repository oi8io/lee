//给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。
//
// 实现 Solution class:
//
//
// Solution(int[] nums) 使用整数数组 nums 初始化对象
// int[] reset() 重设数组到它的初始状态并返回
// int[] shuffle() 返回数组随机打乱后的结果
//
//
//
//
// 示例 1：
//
//
//输入
//["Solution", "shuffle", "reset", "shuffle"]
//[[[1, 2, 3]], [], [], []]
//输出
//[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
//
//解释
//Solution solution = new Solution([1, 2, 3]);
//solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3,
//1, 2]
//solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
//solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 50
// -10⁶ <= nums[i] <= 10⁶
// nums 中的所有元素都是 唯一的
// 最多可以调用 10⁴ 次 reset 和 shuffle
//
// 👍 286 👎 0

package cn

import (
	"math/rand"
)

//leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	r []int
	d []int
	n int
}

func Constructor(nums []int) Solution {
	r := make([]int, len(nums))
	copy(r, nums)
	return Solution{r: nums, d: r, n: len(nums)}
}

func (this *Solution) Reset() []int {
	return this.r
}

func (this *Solution) Shuffle() []int {
	return this.Shuffle2()
	return this.Shuffle1()
}
func (this *Solution) Swap(i, j int) {
	this.d[i], this.d[j] = this.d[j], this.d[i]
}
func (this *Solution) Shuffle2() []int {
	for i := this.n; i >= 1; i-- {
		r := rand.Intn(i)
		this.Swap(i-1, r)
	}
	return this.d
}
func (this *Solution) Shuffle1() []int {
	rand.Shuffle(this.n, func(i, j int) {
		this.d[i], this.d[j] = this.d[j], this.d[i]
	})
	return this.d
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
//leetcode submit region end(Prohibit modification and deletion)
