package main

import (
	"strings"
)

// Tags:
// Trie
func replaceWords(dict []string, sentence string) string {
	trie := NewTrie()
	for _, w := range dict {
		trie.Insert(w)
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if res := trie.Search(word); res != "" {
			words[i] = res
		}
	}
	return strings.Join(words, " ")
}

type Trie struct {
	child  [26]*Trie
	prefix string
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	trie := t
	for i := range word {
		idx := int(word[i] - 'a')
		if trie.child[idx] == nil {
			trie.child[idx] = NewTrie()
		}
		trie = trie.child[idx]
	}
	trie.prefix = word
}

func (t *Trie) Search(word string) string {
	trie := t
	for i := range word {
		idx := int(word[i] - 'a')
		if trie.child[idx] == nil {
			break
		}
		trie = trie.child[idx]
		if trie.prefix != "" {
			break
		}
	}
	return trie.prefix
}

// Prefix Hash
func replaceWords2(dict []string, sentence string) string {
	hasPrefix := make(map[string]bool, len(dict))
	hasLength := make([]bool, 101)
	for i := range dict {
		hasPrefix[dict[i]] = true
		hasLength[len(dict[i])] = true
	}
	lengths := make([]int, 101)
	var index int
	for i, ok := range hasLength {
		if ok {
			lengths[index] = i
			index++
		}
	}
	lengths = lengths[:index]

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 0; j < len(lengths) && lengths[j] < len(word); j++ {
			if hasPrefix[word[:lengths[j]]] {
				words[i] = word[:lengths[j]]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
