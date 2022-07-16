package main

// Difficulty:
// Hash

// Tags:
// Trie

type WordFilter struct {
	t *trie
}

func Constructor(words []string) WordFilter {
	t := &trie{}
	for idx, word := range words {
		for i := range word {
			t.insert(word[i:]+"{"+word, idx)
		}
	}
	return WordFilter{t}
}

func (this *WordFilter) F(pref string, suff string) int {
	return this.t.query(suff + "{" + pref)
}

type trie struct {
	son [27]*trie
	idx int
}

func (t *trie) insert(str string, idx int) {
	for _, c := range str {
		i := int(c - 'a')
		if t.son[i] == nil {
			t.son[i] = &trie{}
		}
		t = t.son[i]
		t.idx = idx
	}
}

func (t *trie) query(str string) int {
	for _, c := range str {
		i := int(c - 'a')
		if t.son[i] == nil {
			return -1
		}
		t = t.son[i]
	}
	return t.idx
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
