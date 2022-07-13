package main

import "strings"

// Difficulty:
// Medium

// Tags:
// Trie

func replaceWords(dictionary []string, sentence string) string {
	t := &trie{}
	for _, word := range dictionary {
		t.insert(word)
	}
	strs := strings.Split(sentence, " ")
	for i, str := range strs {
		if o := t.query(str); o != "" {
			strs[i] = o
		}
	}
	return strings.Join(strs, " ")
}

type trie struct {
	child [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	for _, c := range word {
		i := int(c - 'a')
		if t.child[i] == nil {
			t.child[i] = &trie{}
		}
		t = t.child[i]
	}
	t.isEnd = true
}

func (t *trie) query(word string) string {
	for k, c := range word {
		i := int(c - 'a')
		if t.child[i] == nil {
			break
		}
		if t = t.child[i]; t.isEnd {
			return word[:k+1]
		}
	}
	return ""
}
