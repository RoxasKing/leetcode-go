package main

// Tags:
// Important!!

// Trie + Bit Manipulation
func countPairs(nums []int, low int, high int) int {
	trie := &Trie{}
	out := 0
	for _, num := range nums {
		out += trie.Lower(num, high+1) - trie.Lower(num, low)
		trie.Insert(num)
	}
	return out
}

type Trie struct {
	child [2]*Trie
	count int
}

func (t *Trie) Insert(num int) {
	node := t
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if node.child[idx] == nil {
			node.child[idx] = &Trie{}
		}
		node = node.child[idx]
		node.count++
	}
}

func (t *Trie) Lower(num, low int) int {
	out := 0
	node := t
	for i := 31; i >= 0 && node != nil; i-- {
		lowIdx := (low >> i) & 1
		numIdx := (num >> i) & 1
		if lowIdx == 1 {
			if node.child[numIdx] != nil {
				out += node.child[numIdx].count
			}
			node = node.child[1-numIdx]
		} else {
			node = node.child[numIdx]
		}
	}
	return out
}
