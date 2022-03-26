package questions

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics 数组 回溯 👍 988 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 构建树，将其分解未子问题
// 如11122
// 第一层有两种选择（1）1122 和（2）1112如下图。右边同理
/**
                                                           ┌────────┐
                                          ┌────────────────│ 11122  │────────────────────────┐
                                          │                └────────┘                        │
                                          │                                                  │
                                          ▼                                                  │
                                    ┌──────────┐                                             ▼
                    ┌───────────────│ [1]1122  │────────────────┐                       ┌────────┐
                    │               └──────────┘                │                     ┌─│[2]1112 │────┐
                    │                                           │                     │ └────────┘    │
                    │                                           │                     ▼               ▼
                    ▼                                           ▼                ┌────────┐      ┌────────┐
              ┌──────────┐                                ┌──────────┐           │        │      │        │
      ┌───────│ [11]122  │─┐                           ┌──│ [12]112  │────┐      └────────┘      └────────┘
      │       └──────────┘ │                           │  └──────────┘    │
      │                    │                           │                  │
      ▼                    ▼                           ▼                  ▼
┌──────────┐         ┌──────────┐                ┌──────────┐       ┌──────────┐
│ [111]22  │      ┌──│ [112]12  │─┐            ┌─│ [121]12  │─┐     │ [122]11  │
└──────────┘      │  └──────────┘ │            │ └──────────┘ │     └──────────┘
                  ▼               ▼            ▼              ▼
            ┌──────────┐    ┌──────────┐ ┌──────────┐   ┌──────────┐
            │ [1121]2  │    │ [1122]1  │ │ [1211]2  │   │ [1212]1  │
            └──────────┘    └──────────┘ └──────────┘   └──────────┘
*/
func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	n := len(nums)
	var current [][]int
	//fmt.Println(nums)
	var exists = make(map[int]bool)
	for i := 0; i < n; i++ {
		if _, ok := exists[nums[i]]; ok {
			//fmt.Println("exist:", nums[i])
			continue
		}
		exists[nums[i]] = true
		var nnums []int
		nnums = append(nnums, nums[0:i]...)
		nnums = append(nnums, nums[i+1:]...)
		rows := permuteUnique(nnums)
		for _, row := range rows {
			row = append(row, nums[i])
			current = append(current, row)
		}
	}
	return current
}

//leetcode submit region end(Prohibit modification and deletion)
