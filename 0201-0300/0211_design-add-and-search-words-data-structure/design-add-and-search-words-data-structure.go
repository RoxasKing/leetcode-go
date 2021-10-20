package main

// Difficulty:
// Medium

// Tags:
// Trie
// DFS

type WordDictionary struct {
	t *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{t: &Trie{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.t.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.t, word, 0)
}

type Trie struct {
	next [26]*Trie
	hasW bool
}

func (t *Trie) Insert(word string) {
	for i := range word {
		x := word[i] - 'a'
		if t.next[x] == nil {
			t.next[x] = &Trie{}
		}
		t = t.next[x]
	}
	t.hasW = true
}

func dfs(t *Trie, word string, i int) bool {
	if i == len(word) {
		return t.hasW
	}
	if word[i] != '.' {
		t = t.next[word[i]-'a']
		if t == nil {
			return false
		}
		return dfs(t, word, i+1)
	}
	for j := 0; j < 26; j++ {
		if t.next[j] != nil && dfs(t.next[j], word, i+1) {
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
