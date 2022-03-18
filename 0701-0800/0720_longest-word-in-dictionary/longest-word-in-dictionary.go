package main

// Difficulty:
// Easy

// Tags:
// Trie

func longestWord(words []string) string {
	t := &trie{}
	for _, w := range words {
		t.insert(w)
	}
	out := ""
	for _, w := range words {
		if t.valid(w) && (len(w) > len(out) || len(w) == len(out) && w < out) {
			out = w
		}
	}
	return out
}

type trie struct {
	child [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	for k := range word {
		i := int(word[k] - 'a')
		if t.child[i] == nil {
			t.child[i] = &trie{}
		}
		t = t.child[i]
	}
	t.isEnd = true
}

func (t *trie) valid(word string) bool {
	for k := range word {
		i := int(word[k] - 'a')
		t = t.child[i]
		if !t.isEnd {
			return false
		}
	}
	return true
}
