package main

// Difficulty:
// Medium

// Tags:
// Trie

type MapSum struct {
	t *Trie
}

func Constructor() MapSum {
	return MapSum{t: &Trie{}}
}

func (this *MapSum) Insert(key string, val int) {
	this.t.insert(key, val)
}

func (this *MapSum) Sum(prefix string) int {
	return this.t.getSum(prefix)
}

type Trie struct {
	next     [26]*Trie
	hasK     bool
	val, sum int
}

func (t *Trie) getSum(prefix string) int {
	for i := range prefix {
		if t == nil {
			return 0
		}
		k := prefix[i] - 'a'
		if t.next[k] == nil {
			return 0
		}
		t = t.next[k]
	}
	return t.sum
}

func (t *Trie) insert(key string, val int) {
	kVal := val
	if ok, v := t.getVal(key); ok {
		val -= v
	}
	for i := range key {
		t.sum += val
		k := key[i] - 'a'
		if t.next[k] == nil {
			t.next[k] = &Trie{}
		}
		t = t.next[k]
	}
	t.hasK = true
	t.val = kVal
	t.sum += val
}

func (t *Trie) getVal(key string) (bool, int) {
	for i := range key {
		if t == nil {
			return false, 0
		}
		k := key[i] - 'a'
		if t.next[k] == nil {
			return false, 0
		}
		t = t.next[k]
	}
	return t.hasK, t.val
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
