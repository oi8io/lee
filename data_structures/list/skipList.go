package list

type SkipListNode struct {
	backward *SkipListNode
	next     []level
	key      string
	val      interface{}
}

type level struct {
	span  int
	level int
	next  SkipListNode
}

type skipList struct {
	head, tail *SkipListNode
	len        int
	maxLevel   int
}

type SkipListInf interface {
	Insert(key int, val interface{})
	Find(key int) interface{}
	Delete(key int) interface{}
	Update(key int, val interface{}) interface{}
	Range(min, max int)
}

func New() *skipList {
	head, tail := &SkipListNode{}, &SkipListNode{}
	tail.backward = head
	return &skipList{maxLevel: 32}
}
