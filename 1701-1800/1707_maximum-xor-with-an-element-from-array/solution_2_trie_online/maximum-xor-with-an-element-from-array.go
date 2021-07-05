package main

// Tags:
// Trie
// Online Algorithm
// Bit Manipulation

func maximizeXor(nums []int, queries [][]int) []int {
	t := MakeTrie()
	for _, num := range nums {
		t.Insert(num)
	}

	n := len(queries)
	out := make([]int, n)

	for i, q := range queries {
		out[i] = t.Query(q[0], q[1])
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Trie struct {
	child [2]*Trie
	min   int
}

func MakeTrie() *Trie {
	return &Trie{min: 1<<31 - 1}
}

func (t *Trie) Insert(num int) {
	for b := 31; b >= 0; b-- {
		i := (num >> b) & 1
		if t.child[i] == nil {
			t.child[i] = MakeTrie()
		}
		t = t.child[i]
		t.min = Min(t.min, num)
	}
}

func (t *Trie) Query(x, m int) int {
	out := 0
	for b := 31; b >= 0; b-- {
		i := (x >> b) & 1
		if t.child[1-i] != nil && t.child[1-i].min <= m {
			out |= 1 << b
			t = t.child[1-i]
		} else if t.child[i].min <= m {
			t = t.child[i]
		} else {
			return -1
		}
	}
	return out
}
