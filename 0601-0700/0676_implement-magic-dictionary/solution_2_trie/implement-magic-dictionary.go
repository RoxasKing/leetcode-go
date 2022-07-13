package main

// Difficulty:
// Medium

// Tags:
// Trie

type MagicDictionary struct {
	t *trie
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{t: &trie{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, w := range dictionary {
		this.t.insert(w)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	q := []*triePair{{t: this.t, changed: false}}
	for i := range searchWord {
		idx := int(searchWord[i] - 'a')
		size := len(q)
		for _, p := range q {
			if p.t.child[idx] != nil {
				q = append(q, &triePair{t: p.t.child[idx], changed: p.changed})
			}
			if !p.changed {
				for i := 0; i < 26; i++ {
					if i == idx || p.t.child[i] == nil {
						continue
					}
					q = append(q, &triePair{t: p.t.child[i], changed: true})
				}
			}
		}
		q = q[size:]
	}

	for _, p := range q {
		if p.t.isEnd && p.changed {
			return true
		}
	}
	return false
}

type triePair struct {
	t       *trie
	changed bool
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

type trie struct {
	child [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	for i := range word {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			t.child[idx] = &trie{}
		}
		t = t.child[idx]
	}
	t.isEnd = true
}
