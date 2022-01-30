package main

// Difficulty:
// Medium

// Tags:
// Trie
// DFS

type WordDictionary struct {
	t *trie
}

func Constructor() WordDictionary {
	return WordDictionary{t: &trie{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.t.insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.t.query(word, 0)
}

type trie struct {
	child [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	for i := range word {
		x := word[i] - 'a'
		if t.child[x] == nil {
			t.child[x] = &trie{}
		}
		t = t.child[x]
	}
	t.isEnd = true
}

func (t *trie) query(word string, i int) bool {
	if i == len(word) {
		return t.isEnd
	}
	if word[i] != '.' {
		t = t.child[word[i]-'a']
		if t == nil {
			return false
		}
		return t.query(word, i+1)
	}
	for j := 0; j < 26; j++ {
		if t.child[j] != nil && t.child[j].query(word, i+1) {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
