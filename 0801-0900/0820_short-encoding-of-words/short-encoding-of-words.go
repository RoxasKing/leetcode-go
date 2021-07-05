package main

import "sort"

// Hash
func minimumLengthEncoding(words []string) int {
	mark := make(map[string]struct{})
	for _, word := range words {
		mark[word] = struct{}{}
	}
	for _, word := range words {
		for word != "" {
			word = word[1:]
			delete(mark, word)
		}
	}
	out := 0
	for w := range mark {
		out += len(w) + 1
	}
	return out
}

// Trie
func minimumLengthEncoding2(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	out := 0
	visited := make(map[string]bool)
	for _, word := range words {
		if visited[word] {
			continue
		}
		visited[word] = true
		if trie.TouchEnd(word) {
			out += len(word) + 1
		}
	}
	return out
}

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
	}
}

func (trie *Trie) Insert(word string) {
	t := trie
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			t.child[idx] = NewTrie()
		}
		t = t.child[idx]
		t.isEnd = false
	}
	t.isEnd = true
}

func (trie *Trie) TouchEnd(word string) bool {
	t := trie
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			break
		}
		t = t.child[idx]
	}
	return t.isEnd
}
