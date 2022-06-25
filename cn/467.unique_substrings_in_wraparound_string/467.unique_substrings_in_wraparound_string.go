//把字符串 s 看作是 “abcdefghijklmnopqrstuvwxyz” 的无限环绕字符串，所以 s 看起来是这样的：
//
//
// "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...." .
//
//
// 现在给定另一个字符串 p 。返回 s 中 唯一 的 p 的 非空子串 的数量 。
//
//
//
// 示例 1:
//
//
//输入: p = "a"
//输出: 1
//解释: 字符串 s 中只有一个"a"子字符。
//
//
// 示例 2:
//
//
//输入: p = "cac"
//输出: 2
//解释: 字符串 s 中的字符串“cac”只有两个子串“a”、“c”。.
//
//
// 示例 3:
//
//
//输入: p = "zab"
//输出: 6
//解释: 在字符串 s 中有六个子串“z”、“a”、“b”、“za”、“ab”、“zab”。
//
//
//
//
// 提示:
//
//
// 1 <= p.length <= 10⁵
// p 由小写英文字母构成
//
// 👍 307 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findSubstringInWraproundString(p string) int {
	//zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy
	//	abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz
	//	vwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstu
	//	uvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrst
	//z->y
	// b->a
	// abcd a ,b ,c ,d ab,bc,cd ,abc,bcd, abcd

	// abcd cde  a ,b ,c ,d ab,bc,cd ,abc,bcd, abcd
	//var count 'a'+iint
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var tmp = 1
	var xmap = make(map[byte]int)
	xmap[p[0]-'a'] = 1
	for i := 1; i < len(p); i++ {
		if (p[i]-p[i-1]+26)%26 == 1 {
			tmp++
		} else {
			tmp = 1
		}
		xmap[p[i]-'a'] = max(xmap[p[i]-'a'], tmp)
	}
	var ans int
	for i := 0; i < 26; i++ {
		ans += xmap[byte(i)]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)