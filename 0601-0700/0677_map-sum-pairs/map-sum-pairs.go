package main

// Tags:
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
