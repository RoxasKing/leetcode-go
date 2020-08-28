package main

/*
  实现一个 MapSum 类里的两个方法，insert 和 sum。

  对于方法 insert，你将得到一对（字符串，整数）的键值对。字符串表示键，整数表示值。如果键已经存在，那么原来的键值对将被替代成新的键值对。

  对于方法 sum，你将得到一个表示前缀的字符串，你需要返回所有以该前缀开头的键的值的总和。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/map-sum-pairs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MapSum struct {
	m     [26]*MapSum
	sum   int
	isEnd bool
}

/** Initialize your data structure here. */
func NewMapSum() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	for i := range key {
		k := key[i] - 'a'
		if this.m[k] == nil {
			this.m[k] = new(MapSum)
		}
		this = this.m[k]
	}
	this.isEnd = true
	this.sum = val
}

func (this *MapSum) Sum(prefix string) int {
	for i := range prefix {
		k := prefix[i] - 'a'
		if this.m[k] == nil {
			return 0
		}
		this = this.m[k]
	}
	var out int
	q := []*MapSum{this}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		if cur.isEnd {
			out += cur.sum
		}
		for i := range cur.m {
			if cur.m[i] != nil {
				q = append(q, cur.m[i])
			}
		}
	}
	return out
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
