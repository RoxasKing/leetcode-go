package leetcode

/*
  实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

  说明:
    你可以假设所有的输入都是由小写字母 a-z 构成的。
    保证所有输入均为非空字符串。
*/

type Trie struct {
	trie  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func NewTrie() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, c := range word {
		if this.trie[c-'a'] == nil {
			this.trie[c-'a'] = new(Trie)
		}
		this = this.trie[c-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, c := range word {
		if this.trie[c-'a'] == nil {
			return false
		}
		this = this.trie[c-'a']
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if this.trie[c-'a'] == nil {
			return false
		}
		this = this.trie[c-'a']
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
