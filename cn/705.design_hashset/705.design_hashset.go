//不使用任何内建的哈希表库设计一个哈希集合（HashSet）。
//
// 实现 MyHashSet 类：
//
//
// void add(key) 向哈希集合中插入值 key 。
// bool contains(key) 返回哈希集合中是否存在这个值 key 。
// void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
//
//
//
// 示例：
//
//
//输入：
//["MyHashSet", "add", "add", "contains", "contains", "add", "contains",
//"remove", "contains"]
//[[], [1], [2], [1], [3], [2], [2], [2], [2]]
//输出：
//[null, null, null, true, false, null, true, null, false]
//
//解释：
//MyHashSet myHashSet = new MyHashSet();
//myHashSet.add(1);      // set = [1]
//myHashSet.add(2);      // set = [1, 2]
//myHashSet.contains(1); // 返回 True
//myHashSet.contains(3); // 返回 False ，（未找到）
//myHashSet.add(2);      // set = [1, 2]
//myHashSet.contains(2); // 返回 True
//myHashSet.remove(2);   // set = [1]
//myHashSet.contains(2); // 返回 False ，（已移除）
//
//
//
// 提示：
//
//
// 0 <= key <= 10⁶
// 最多调用 10⁴ 次 add、remove 和 contains
//
// 👍 242 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
type MyHashSet struct {
	data [][]int
	n    int
	cnt  int
}

func Constructor() MyHashSet {
	return MyHashSet{make([][]int, 1<<8), 8, 0}
}

func (s *MyHashSet) Add(key int) {
	if s.Contains(key) {
		return
	}
	bucket := s.GetBucket(key)
	s.cnt++
	s.data[bucket] = append(s.data[bucket], key)
	if s.cnt/1<<s.n > 4 { // todo 扩容

	}
}

func (s *MyHashSet) GetBucket(key int) int {
	bucket := key & (1<<s.n - 1)
	return bucket
}
func (s *MyHashSet) Remove(key int) {
	b := s.GetBucket(key)
	a := s.data[b]
	n := len(a)
	if n == 0 {
		return
	}
	i := 0
	for ; i < len(a); i++ {
		if a[i] == key {
			break
		}
	}
	if i >= len(a)-1 {
		a = a[:i]
	} else {
		a = append(a[:i], a[i+1:]...)
	}
	s.data[b] = a
}

func (s *MyHashSet) Contains(key int) bool {
	b := s.GetBucket(key)
	a := s.data[b]
	for i := 0; i < len(a); i++ {
		if a[i] == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
//leetcode submit region end(Prohibit modification and deletion)
