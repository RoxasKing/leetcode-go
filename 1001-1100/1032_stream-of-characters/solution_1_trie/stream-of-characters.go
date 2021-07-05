package main

// Tags:
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
