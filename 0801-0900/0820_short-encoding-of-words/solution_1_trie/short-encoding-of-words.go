package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Trie

func minimumLengthEncoding(words []string) int {
	t := &trie{}
	sort.Slice(words, func(i, j int) bool { return len(words[i]) > len(words[j]) })
	o := 0
	for _, word := range words {
		if t.insert(word) {
			o += len(word) + 1
		}
	}
	return o
}

type trie struct{ child [26]*trie }

func (t *trie) insert(word string) bool {
	add := false
	for i := len(word) - 1; i >= 0; i-- {
		k := int(word[i] - 'a')
		if t.child[k] == nil {
			t.child[k] = &trie{}
			add = true
		}
		t = t.child[k]
	}
	return add
}
