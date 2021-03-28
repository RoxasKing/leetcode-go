package main

/*
  实现一个 MapSum 类里的两个方法，insert 和 sum。

  对于方法 insert，你将得到一对（字符串，整数）的键值对。字符串表示键，整数表示值。如果键已经存在，那么原来的键值对将被替代成新的键值对。

  对于方法 sum，你将得到一个表示前缀的字符串，你需要返回所有以该前缀开头的键的值的总和。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/map-sum-pairs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie + Queue
type MapSum struct {
	trie *Trie
}

type Trie struct {
	child [26]*Trie
	num   int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		trie: &Trie{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	t := this.trie
	for i := range key {
		idx := int(key[i] - 'a')
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}
	t.num = val
}

func (this *MapSum) Sum(prefix string) int {
	t := this.trie
	for i := range prefix {
		idx := int(prefix[i] - 'a')
		if t.child[idx] == nil {
			return 0
		}
		t = t.child[idx]
	}

	sum := 0
	q := []*Trie{t}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		sum += t.num
		for i := 0; i < 26; i++ {
			if t.child[i] != nil {
				q = append(q, t.child[i])
			}
		}
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
