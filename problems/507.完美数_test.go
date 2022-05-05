/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 *
 * https://leetcode-cn.com/problems/perfect-number/description/
 *
 * algorithms
 * Easy (48.94%)
 * Likes:    179
 * Dislikes: 0
 * Total Accepted:    64.7K
 * Total Submissions: 132.1K
 * Testcase Example:  '28'
 *
 * 对于一个 正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。
 *
 * 给定一个 整数 n， 如果是完美数，返回 true；否则返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = 28
 * 输出：true
 * 解释：28 = 1 + 2 + 4 + 7 + 14
 * 1, 2, 4, 7, 和 14 是 28 的所有正因子。
 *
 * 示例 2：
 *
 *
 * 输入：num = 7
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num <= 10^8
 *
 *
 */

package problems

import (
	"fmt"
	"testing"
	"time"
)

func Test_checkPerfectNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{18}, false},
		{"", args{28}, true},
		{"", args{99999995}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{2016}, false},
		{"", args{8128}, true},
		{"", args{496}, true},
		{"", args{28}, true},
		{"", args{6}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPerfectNumber(t *testing.T) {
	var nums []int
	before := time.Now()
	for i := 0; i <= 10000000; i++ {
		got := checkPerfectNumber(i)
		if got {
			nums = append(nums, i)
		}
		if i > 0 && i%10000 == 0 {
			now := time.Now()
			after := time.Now().Sub(before)
			fmt.Println(i, after.Milliseconds())
			before = now
		}
	}
	fmt.Println(nums)
}
