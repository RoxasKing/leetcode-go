package main

// Tags:
// Aho–Corasick Algorithm + Trie
type StreamChecker struct {
	t *Trie
}

func Constructor(words []string) StreamChecker {
	t := &Trie{}
	for _, w := range words {
		t.Insert(w)
	}
	t.Build()

	return StreamChecker{t: t}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.t = this.t.child[letter-'a']
	return this.t.isEnd
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

type Trie struct {
	child [26]*Trie
	isEnd bool
	fail  *Trie
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

func (t *Trie) Build() {
	t.fail = t
	q := []*Trie{}

	for i := 0; i < 26; i++ {
		if t.child[i] != nil {
			t.child[i].fail = t
			q = append(q, t.child[i])
		} else {
			t.child[i] = t
		}
	}

	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		t.isEnd = t.isEnd || t.fail.isEnd
		for i := 0; i < 26; i++ {
			if t.child[i] != nil {
				t.child[i].fail = t.fail.child[i]
				q = append(q, t.child[i])
			} else {
				t.child[i] = t.fail.child[i]
			}
		}
	}
}
