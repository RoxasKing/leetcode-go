package main

// Tags:
// Trie

func longestWord(words []string) string {
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}

	var out string
	for _, word := range words {
		if t.IsValid(word) && (len(word) > len(out) || len(word) == len(out) && word < out) {
			out = word
		}
	}
	return out
}

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func (t *Trie) Insert(word string) {
	for i := range word {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}
	t.isEnd = true
}

func (t *Trie) IsValid(word string) bool {
	for i := range word {
		idx := int(word[i] - 'a')
		t = t.child[idx]
		if !t.isEnd {
			return false
		}
	}
	return true
}
