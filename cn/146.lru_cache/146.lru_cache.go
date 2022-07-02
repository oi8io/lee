//请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
//
// 实现 LRUCache 类：
//
//
//
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
//key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
//
//
// 示例：
//
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 10⁵
// 最多调用 2 * 10⁵ 次 get 和 put
//
// 👍 2251 👎 0

package cn

import "container/list"

//leetcode submit region begin(Prohibit modification and deletion)

type LRUCache struct {
	m   map[int]*list.Element
	l   *list.List
	len int
	cap int
}

type Store struct {
	Key int
	Val int
}

func Constructor(capacity int) LRUCache {
	l := list.New()
	c := LRUCache{
		m:   make(map[int]*list.Element),
		l:   l,
		len: 0,
		cap: capacity,
	}
	return c
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.m[key]; !ok {
		return -1
	} else {
		c.l.MoveToFront(v)
		return v.Value.(Store).Val
	}
}

func (c *LRUCache) Put(key int, value int) {
	s := Store{key, value}
	if v, ok := c.m[key]; ok {
		v.Value = s
		c.l.MoveToFront(v)
	} else {
		if c.len == c.cap { //full
			last := c.l.Back()
			delete(c.m, last.Value.(Store).Key)
			c.l.Remove(last)
		} else {
			c.len++
		}
		n := c.l.PushFront(s)
		c.m[key] = n
	}
}

// 方案1
/*type LRUCache struct {
	m    map[int]*DoubleLinkedList
	head *DoubleLinkedList
	end  *DoubleLinkedList
	len  int
	cap  int
}
type DoubleLinkedList struct {
	Key  int
	Val  int
	Prev *DoubleLinkedList
	Next *DoubleLinkedList
}

func Constructor(capacity int) LRUCache {
	h, e := &DoubleLinkedList{}, &DoubleLinkedList{}
	h.Next = e
	e.Prev = h
	c := LRUCache{
		m:    make(map[int]*DoubleLinkedList),
		head: h,
		end:  e,
		len:  0,
		cap:  capacity,
	}
	return c
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.m[key]; !ok {
		return -1
	} else {
		// 前面的指针
		if v.Prev != nil {
			v.Prev.Next = v.Next
		}
		//后面的指针
		if v.Next != nil {
			v.Next.Prev = v.Prev
		}
		//替换头
		v.Next = c.head.Next
		c.head.Next.Prev = v
		c.head.Next = v
		v.Prev = c.head
		return v.Val
	}
}

func (c *LRUCache) Put(key int, value int) {
	n := &DoubleLinkedList{Key: key, Val: value}
	if v, ok := c.m[key]; ok {
		v.Prev.Next = v.Next //删除当前节点
		v.Next.Prev = v.Prev
	} else {
		if c.len == c.cap { //full
			//淘汰
			delete(c.m, c.end.Prev.Key)
			c.end.Prev = c.end.Prev.Prev
			c.end.Prev.Next = c.end
			//前进
		} else {
			c.len++
		}
	}
	n.Next = c.head.Next
	n.Next.Prev = n
	c.head.Next = n
	n.Prev = c.head
	c.m[key] = n
}
*/
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
