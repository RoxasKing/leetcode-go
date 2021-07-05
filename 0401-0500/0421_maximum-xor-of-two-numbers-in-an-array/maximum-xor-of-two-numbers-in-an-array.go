package main

// Tags:
// Trie
func findMaximumXOR(nums []int) int {
	out := 0
	t := NewTrie()
	for _, num := range nums {
		if res := t.Query(num); res > out {
			out = res
		}
		t.Insert(num)
	}
	return out
}

type Trie struct {
	child [2]*Trie
}

func NewTrie() *Trie {
	return &Trie{child: [2]*Trie{}}
}

func (t *Trie) Insert(num int) {
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if t.child[idx] == nil {
			t.child[idx] = NewTrie()
		}
		t = t.child[idx]
	}
}

func (t *Trie) Query(num int) int {
	out := 0
	for i := 31; i >= 0; i-- {
		if t == nil {
			return 0
		}
		idx := (num >> i) & 1
		if t.child[1-idx] != nil {
			out |= 1 << i
			t = t.child[1-idx]
		} else {
			t = t.child[idx]
		}
	}
	return out
}
