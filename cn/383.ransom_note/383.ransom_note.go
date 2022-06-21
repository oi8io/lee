//给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//
// 如果可以，返回 true ；否则返回 false 。
//
// magazine 中的每个字符只能在 ransomNote 中使用一次。
//
//
//
// 示例 1：
//
//
//输入：ransomNote = "a", magazine = "b"
//输出：false
//
//
// 示例 2：
//
//
//输入：ransomNote = "aa", magazine = "ab"
//输出：false
//
//
// 示例 3：
//
//
//输入：ransomNote = "aa", magazine = "aab"
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= ransomNote.length, magazine.length <= 10⁵
// ransomNote 和 magazine 由小写英文字母组成
//
// 👍 348 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	var m = make([]int, 26, 26)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		m[ransomNote[i]-'a']--
		if m[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
func canConstruct1(ransomNote string, magazine string) bool {
	var m = make([]int, 26, 26)
	for i := 0; i < len(ransomNote); i++ {
		m[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if m[i] > 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
