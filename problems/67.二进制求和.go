/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

package problems

// @lc code=start
func addBinary(a string, b string) string {
	la, lb := len(a)-1, len(b)-1
	l := lb
	if la > lb {
		l = la
	}
	var carry byte = 48
	var s []byte
	for i := 0; i <= l; i++ {
		var ai, bi byte = 48, 48
		if lb-i >= 0 {
			bi = b[lb-i]
		}
		if la-i >= 0 {
			ai = a[la-i]
		}
		var char byte
		char, carry = cal(ai, bi, carry)
		s = append([]byte{char}, s...)
	}
	if carry == 49 {
		s = append([]byte{49}, s...)
	}
	return string(s)
}

func cal(a, b, carry byte) (byte, byte) {
	var char byte
	switch carry + b + a - 48*3 {
	case 0:
		carry = 48
		char = 48
	case 1:
		carry = 48
		char = 49
	case 2:
		carry = 49
		char = 48
	case 3:
		carry = 49
		char = 49
	}
	return char, carry
}

// @lc code=end
