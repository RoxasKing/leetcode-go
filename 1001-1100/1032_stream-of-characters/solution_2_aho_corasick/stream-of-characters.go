package main

/*
  Implement the StreamChecker class as follows:
    1. StreamChecker(words): Constructor, init the data structure with the given words.
    2. query(letter): returns true if and only if for some k >= 1, the last k characters queried (in order from oldest to newest, including this letter just queried) spell one of the words in the given list.

  Example:
    StreamChecker streamChecker = new StreamChecker(["cd","f","kl"]); // init the dictionary.
    streamChecker.query('a');          // return false
    streamChecker.query('b');          // return false
    streamChecker.query('c');          // return false
    streamChecker.query('d');          // return true, because 'cd' is in the wordlist
    streamChecker.query('e');          // return false
    streamChecker.query('f');          // return true, because 'f' is in the wordlist
    streamChecker.query('g');          // return false
    streamChecker.query('h');          // return false
    streamChecker.query('i');          // return false
    streamChecker.query('j');          // return false
    streamChecker.query('k');          // return false
    streamChecker.query('l');          // return true, because 'kl' is in the wordlist

  Note:
    1. 1 <= words.length <= 2000
    2. 1 <= words[i].length <= 2000
    3. Words will only consist of lowercase English letters.
    4. Queries will only consist of lowercase English letters.
    5. The number of queries is at most 40000.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/stream-of-characters
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

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
