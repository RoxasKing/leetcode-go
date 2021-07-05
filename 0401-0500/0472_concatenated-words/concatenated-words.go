package main

import (
	"sort"
)

// Tags:
// Dynamic Programming + Hash
func findAllConcatenatedWordsInADict(words []string) []string {
	sets := make(map[string]bool)
	lens := make(map[int]bool)
	for _, w := range words {
		if w == "" {
			continue
		}
		sets[w] = true
		lens[len(w)] = true
	}
	ls := make([]int, 0, len(lens))
	for l := range lens {
		ls = append(ls, l)
	}
	sort.Ints(ls)
	out := []string{}
	for _, w := range words {
		if w == "" {
			continue
		}
		delete(sets, w)
		n := len(w)
		dp := make([]bool, n+1)
		dp[0] = true
		for i := 1; i <= n; i++ {
			for _, l := range ls {
				if i-l < 0 {
					break
				} else if dp[i-l] && sets[w[i-l:i]] {
					dp[i] = true
					break
				}
			}
		}
		if dp[n] {
			out = append(out, w)
		}
		sets[w] = true
	}
	return out
}

// Dynamic Programming + Trie
func findAllConcatenatedWordsInADict2(words []string) []string {
	t := &Trie{}
	lenSets := make(map[int]bool)
	for _, w := range words {
		if w == "" {
			continue
		}
		t.Insert(w)
		lenSets[len(w)] = true
	}
	lens := make([]int, 0, len(lenSets))
	for l := range lenSets {
		lens = append(lens, l)
	}
	sort.Ints(lens)
	out := []string{}
	for _, w := range words {
		if w == "" {
			continue
		}
		n := len(w)
		dp := make([]bool, n+1)
		dp[0] = true
		for i := 1; i <= n; i++ {
			for _, l := range lens {
				if (i-l == 0 && i < n || i-l > 0) && dp[i-l] && t.Search(w[i-l:i]) {
					dp[i] = true
					break
				} else if i-l < 0 {
					break
				}
			}
		}
		if dp[n] {
			out = append(out, w)
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

func (t *Trie) Search(word string) bool {
	for i := range word {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			return false
		}
		t = t.child[idx]
	}
	return t.isEnd
}
