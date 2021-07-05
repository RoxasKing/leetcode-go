package main

// Tags:
// Trie

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, ch := range word {
		if this.next[ch-'a'] == nil {
			this.next[ch-'a'] = &Trie{}
		}
		this = this.next[ch-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, ch := range word {
		if this.next[ch-'a'] == nil {
			return false
		}
		this = this.next[ch-'a']
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, ch := range prefix {
		if this.next[ch-'a'] == nil {
			return false
		}
		this = this.next[ch-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
