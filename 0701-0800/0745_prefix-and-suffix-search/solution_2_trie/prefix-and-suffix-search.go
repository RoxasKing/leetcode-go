package main

// Difficulty:
// Hash

// Tags:
// Trie

type WordFilter struct {
	trie *Trie
}

func Constructor(words []string) WordFilter {
	t := &Trie{}
	for i, word := range words {
		for j := range word {
			t.Insert(word[j:]+"#"+word, i)
		}
	}
	return WordFilter{t}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	return this.trie.Search(suffix + "#" + prefix)
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */

type Trie struct {
	child  [27]*Trie
	isEnd  bool
	weight int
}

func (t *Trie) Insert(word string, weight int) {
	for i := range word {
		idx := 0
		if word[i] != '#' {
			idx = int(word[i] - 'a' + 1)
		}
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}
	t.isEnd = true
	t.weight = weight
}

func (t *Trie) Search(prefix string) int {
	for i := range prefix {
		idx := 0
		if prefix[i] != '#' {
			idx = int(prefix[i] - 'a' + 1)
		}
		if t.child[idx] == nil {
			return -1
		}
		t = t.child[idx]
	}

	out := -1
	q := []*Trie{t}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		if t.isEnd {
			out = Max(out, t.weight)
		}
		for i := 1; i <= 26; i++ {
			if t.child[i] != nil {
				q = append(q, t.child[i])
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
