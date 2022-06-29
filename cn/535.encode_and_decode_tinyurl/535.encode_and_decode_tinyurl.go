//TinyURL 是一种 URL 简化服务， 比如：当你输入一个 URL https://leetcode.com/problems/design-
//tinyurl 时，它将返回一个简化的URL http://tinyurl.com/4e9iAk 。请你设计一个类来加密与解密 TinyURL 。
//
// 加密和解密算法如何设计和运作是没有限制的，你只需要保证一个 URL 可以被加密成一个 TinyURL ，并且这个 TinyURL 可以用解密方法恢复成原本
//的 URL 。
//
// 实现 Solution 类：
//
//
//
//
// Solution() 初始化 TinyURL 系统对象。
// String encode(String longUrl) 返回 longUrl 对应的 TinyURL 。
// String decode(String shortUrl) 返回 shortUrl 原本的 URL 。题目数据保证给定的 shortUrl 是由同一个系
//统对象加密的。
//
//
//
//
// 示例：
//
//
//输入：url = "https://leetcode.com/problems/design-tinyurl"
//输出："https://leetcode.com/problems/design-tinyurl"
//
//解释：
//Solution obj = new Solution();
//string tiny = obj.encode(url); // 返回加密后得到的 TinyURL 。
//string ans = obj.decode(tiny); // 返回解密后得到的原本的 URL 。
//
//
//
//
// 提示：
//
//
// 1 <= url.length <= 10⁴
// 题目数据保证 url 是一个有效的 URL
//
//
//
// 👍 150 👎 0

package cn

import (
	"math/rand"
	"time"
)

//leetcode submit region begin(Prohibit modification and deletion)
type Codec struct {
	N     int
	store map[string]string
}

func Constructor() Codec {
	return Codec{store: make(map[string]string), N: 56800235584}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	randx := rand.New(rand.NewSource(time.Now().UnixNano()))
	intn := randx.Intn(this.N)
	var x = make([]byte, 0)
	for intn > 0 {
		cur := byte(intn % 62)
		var c byte = '0'
		switch {
		case cur < 10:
			c = cur + '0'
		case cur >= 10 && cur < 36:
			c = cur - 10 + 'a'
		case cur >= 36:
			c = cur - 36 + 'A'
		}
		x = append(x, c)
		intn = intn / 62
	}

	s := string(x)
	this.store[s] = longUrl
	return "https://x.io/" + s
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	s := shortUrl[13:]
	return this.store[s]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
//leetcode submit region end(Prohibit modification and deletion)
