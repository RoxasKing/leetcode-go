package main

// Difficulty:
// Medium

// Tags:
// Trie
// BFS

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
	return this.t.query(word)
}

type trie struct {
	child [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	for i := range word {
		idx := word[i] - 'a'
		if t.child[idx] == nil {
			t.child[idx] = &trie{}
		}
		t = t.child[idx]
	}
	t.isEnd = true
}

func (t *trie) query(word string) bool {
	q := []*trie{t}
	for i := range word {
		n := len(q)
		for _, t := range q {
			if word[i] != '.' {
				if t.child[word[i]-'a'] != nil {
					q = append(q, t.child[word[i]-'a'])
				}
			} else {
				for j := 0; j < 26; j++ {
					if t.child[j] != nil {
						q = append(q, t.child[j])
					}
				}
			}
		}
		q = q[n:]
	}
	for _, t := range q {
		if t.isEnd {
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
