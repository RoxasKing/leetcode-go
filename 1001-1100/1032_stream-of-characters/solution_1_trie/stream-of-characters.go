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

// Trie
type StreamChecker struct {
	trie   *Trie
	limit  int
	buffer []byte
}

func Constructor(words []string) StreamChecker {
	trie := &Trie{}
	limit := 0
	for _, w := range words {
		limit = Max(limit, len(w))
		arr := make([]byte, len(w))
		for i := range w {
			arr[i] = w[i] - 'a'
		}
		reverse(arr)
		trie.Insert(arr)
	}
	return StreamChecker{
		trie:   trie,
		limit:  limit,
		buffer: []byte{},
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.buffer = append(this.buffer, letter-'a')
	if len(this.buffer) > this.limit {
		this.buffer = this.buffer[1:]
	}

	idxs := []byte{}
	for i := len(this.buffer) - 1; i >= 0; i-- {
		idxs = append(idxs, this.buffer[i])
		res := this.trie.Search(idxs)
		if res == -1 {
			break
		} else if res == 1 {
			return true
		}
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(arr []byte) {
	n := len(arr)
	for i := 0; i < n>>1; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func (t *Trie) Insert(idxs []byte) {
	for _, idx := range idxs {
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}
	t.isEnd = true
}

func (t *Trie) Search(idxs []byte) int {
	for _, idx := range idxs {
		if t.child[idx] == nil {
			return -1
		}
		t = t.child[idx]
		if t.isEnd {
			return 1
		}
	}
	return 0
}
