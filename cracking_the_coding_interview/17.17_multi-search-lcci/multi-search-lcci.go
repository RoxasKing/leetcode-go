package main

// Tags:
// Trie
func multiSearch(big string, smalls []string) [][]int {
	t := NewTrie()
	for _, w := range smalls {
		t.Insert(w)
	}
	for i := range big {
		t.AddIdx(big[i:], i)
	}
	n := len(smalls)
	out := make([][]int, n)
	for i, w := range smalls {
		out[i] = t.Query(w)
	}
	return out
}

type Trie struct {
	child [26]*Trie
	isEnd bool
	idxs  []int
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
		isEnd: false,
		idxs:  []int{},
	}
}

func (t *Trie) Insert(word string) {
	for i := range word {
		idx := word[i] - 'a'
		if t.child[idx] == nil {
			t.child[idx] = NewTrie()
		}
		t = t.child[idx]
	}
	t.isEnd = true
}

func (t *Trie) AddIdx(str string, idx int) {
	for i := range str {
		j := str[i] - 'a'
		if t.child[j] == nil {
			break
		}
		t = t.child[j]
		if t.isEnd {
			t.idxs = append(t.idxs, idx)
		}
	}
}

func (t *Trie) Query(word string) []int {
	for i := range word {
		idx := word[i] - 'a'
		if t.child[idx] == nil {
			break
		}
		t = t.child[idx]
	}
	return t.idxs
}
