package main

// Difficulty:
// Medium

// Tags:
// Trie

func findMaximumXOR(nums []int) int {
	out := 0
	t := &trie{}
	for _, num := range nums {
		if res := t.query(num); res > out {
			out = res
		}
		t.insert(num)
	}
	return out
}

type trie struct{ child [2]*trie }

func (t *trie) insert(num int) {
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if t.child[idx] == nil {
			t.child[idx] = &trie{}
		}
		t = t.child[idx]
	}
}

func (t *trie) query(num int) int {
	out := 0
	for i := 31; i >= 0 && t != nil; i-- {
		out <<= 1
		idx := (num >> i) & 1
		if t.child[idx^1] != nil {
			out++
			idx ^= 1
		}
		t = t.child[idx]
	}
	return out
}
