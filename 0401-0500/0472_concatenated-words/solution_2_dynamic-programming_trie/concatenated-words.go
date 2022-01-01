package main

import "sort"

// Difficulty:
// Hard

// Tags
// Dynamic Programming
// Trie

func findAllConcatenatedWordsInADict(words []string) []string {
	t := &Trie{}
	sizeSet := map[int]struct{}{}
	var idx int
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		words[idx] = word
		idx++
		t.Insert(word)
		sizeSet[len(word)] = struct{}{}
	}
	words = words[:idx]
	sizes := make([]int, 0, len(sizeSet))
	for size := range sizeSet {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	var out []string
	for _, word := range words {
		t.SoftDel(word)
		n := len(word)
		f := make([]bool, n+1)
		f[0] = true
		for i := 1; i <= n; i++ {
			for _, size := range sizes {
				if i-size < 0 {
					break
				}
				if f[i-size] && t.Exists(word[i-size:i]) {
					f[i] = true
					break
				}
			}
		}
		if f[n] {
			out = append(out, word)
		}
		t.Insert(word)
	}
	return out
}

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func (t *Trie) Insert(word string) {
	for i := range word {
		idx := word[i] - 'a'
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}
	t.isEnd = true
}

func (t *Trie) Exists(word string) bool {
	for i := range word {
		idx := word[i] - 'a'
		if t.child[idx] == nil {
			return false
		}
		t = t.child[idx]
	}
	return t.isEnd
}

func (t *Trie) SoftDel(word string) {
	for i := range word {
		t = t.child[word[i]-'a']
	}
	t.isEnd = false
}
