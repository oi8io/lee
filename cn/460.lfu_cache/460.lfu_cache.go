//请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
//
// 实现 LFUCache 类：
//
//
// LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
// int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
// void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
//capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
//
//
// 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
//
// 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
// 示例：
//
//
//输入：
//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
//"get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
//输出：
//[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//
//解释：
//// cnt(x) = 键 x 的使用计数
//// ele=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
//LFUCache lfu = new LFUCache(2);
//lfu.put(1, 1);   // ele=[1,_], cnt(1)=1
//lfu.put(2, 2);   // ele=[2,1], cnt(2)=1, cnt(1)=1
//lfu.get(1);      // 返回 1
//                 // ele=[1,2], cnt(2)=1, cnt(1)=2
//lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
//                 // ele=[3,1], cnt(3)=1, cnt(1)=2
//lfu.get(2);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3
//                 // ele=[3,1], cnt(3)=2, cnt(1)=2
//lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
//                 // ele=[4,3], cnt(4)=1, cnt(3)=2
//lfu.get(1);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3
//                 // ele=[3,4], cnt(4)=1, cnt(3)=3
//lfu.get(4);      // 返回 4
//                 // ele=[3,4], cnt(4)=2, cnt(3)=3
//
//
//
// 提示：
//
//
// 0 <= capacity <= 10⁴
// 0 <= key <= 10⁵
// 0 <= value <= 10⁹
// 最多调用 2 * 10⁵ 次 get 和 put 方法
//
// 👍 563 👎 0

package cn

import (
	"container/list"
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
type LFUCache struct { // todo 线段树
	ele  map[int]*list.Element
	list map[int]*list.List
	f    []int
	cap  int
	len  int
}
type Store struct {
	key, val, freq int
}

func Constructor(capacity int) LFUCache {
	c := LFUCache{
		ele:  make(map[int]*list.Element),
		list: make(map[int]*list.List),
		f:    make([]int, 0),
		cap:  capacity,
		len:  0,
	}
	return c
}

// Remove 移除元素
// 1 先找到成绩，再获取链表
// 2 删除元素记录
func (c *LFUCache) Remove(v *list.Element) Store {
	s := lv(v)
	f := s.freq
	l := c.GetList(f)
	delete(c.ele, s.key)
	remove := l.Remove(v)
	if l.Len() == 0 {
		delete(c.list, f) //
		c.deleteFreq(f)   // 删除层
	}
	return remove.(Store)
}
func (c *LFUCache) Get(key int) int {
	if v, ok := c.ele[key]; ok {
		s := c.Remove(v)
		s.freq++
		back := c.GetList(s.freq).PushFront(s)
		c.ele[key] = back
		return s.val
	} else {
		return -1
	}
}

func (c *LFUCache) GetList(freq int) *list.List {
	if v, ok := c.list[freq]; ok {
		return v
	} else {
		c.list[freq] = list.New()
		c.insertFreq(freq)
		return c.list[freq]
	}
}
func (s Store) String() string {
	return fmt.Sprintf("key:[%d]val:[%d]freq:[%d]", s.key, s.val, s.freq)
}

func lv(e *list.Element) Store {
	if e == nil {
		return Store{}
	}
	return e.Value.(Store)
}

func (c *LFUCache) Put(key int, value int) {
	if c.cap == 0 {
		return
	}
	if v, ok := c.ele[key]; ok { //update
		s := c.Remove(v)
		s.freq++
		s.val = value
		back := c.GetList(s.freq).PushFront(s)
		c.ele[key] = back
	} else {
		if c.cap == c.len {
			f := c.minFreq()
			l := c.GetList(f)
			back := l.Back()
			c.Remove(back)
		} else {
			c.len++
		}
		s := Store{key, value, 1}
		front := c.GetList(1).PushFront(s)
		c.ele[key] = front
	}
}

func (c *LFUCache) deleteFreq(n int) {
	freq := c.findFreq(n)
	l := len(c.f)
	k := freq % l
	if c.f[k] != n {
		return
	}
	if freq == 0 {
		c.f = c.f[1:]
	} else if freq == len(c.f) {
		c.f = c.f[:len(c.f)-1]
	} else {
		c.f = append(c.f[:freq], c.f[freq+1:]...)
	}
}
func (c *LFUCache) insertFreq(n int) {
	if len(c.f) == 0 {
		c.f = []int{n}
		return
	}
	freq := c.findFreq(n)
	if freq == 0 {
		c.f = append([]int{n}, c.f...)
	} else if freq == len(c.f) {
		c.f = append(c.f, n)
	} else {
		if c.f[freq] == n {
			return
		}
		nx := append([]int{}, c.f[:freq]...)
		nx = append(nx, n)
		nx = append(nx, c.f[freq:]...)
		c.f = nx
	}
}
func (c *LFUCache) findFreq(n int) int {
	min, max := 0, len(c.f)
	var m int
	for {
		m = (min + max) / 2
		if min >= max-1 {
			break
		}
		if c.f[m] > n {
			max = m
		} else {
			if m == min {
				min++
			} else {
				min = m
			}
		}
	}
	if c.f[m] >= n {
		return m
	} else {
		return m + 1
	}
}

func (c *LFUCache) minFreq() int {
	return c.f[0]
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
