package main

/*
  实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

  说明:
    你可以假设所有的输入都是由小写字母 a-z 构成的。
    保证所有输入均为非空字符串。
*/

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
	for _, letter := range word {
		if this.next[letter-'a'] == nil {
			this.next[letter-'a'] = new(Trie)
		}
		this = this.next[letter-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, letter := range word {
		if this.next[letter-'a'] == nil {
			return false
		}
		this = this.next[letter-'a']
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, letter := range prefix {
		if this.next[letter-'a'] == nil {
			return false
		}
		this = this.next[letter-'a']
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
