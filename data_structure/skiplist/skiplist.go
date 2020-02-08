package skiplist

import "math"

const (
	MAX_LEVEL = 16
)

type SkipList struct {
	head   *skipListNode
	level  int
	length int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:   newSkipListNode(0, math.MinInt32, MAX_LEVEL),
		level:  1,
		length: 0,
	}
}

type skipListNode struct {
	v       interface{}
	score   int
	level   int
	forward []*skipListNode
}

func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{
		v:       v,
		score:   score,
		level:   level,
		forward: make([]*skipListNode, level, level),
	}
}

func (sl *SkipList) Length() int {
	return sl.length
}

func (sl *SkipList) Level() int {
	return sl.level
}

