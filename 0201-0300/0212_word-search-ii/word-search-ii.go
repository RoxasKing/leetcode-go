package main

// Tags:
// Trie
// Backtracking

func findWords(board [][]byte, words []string) []string {
	t := &Trie{}
	for _, w := range words {
		t.Insert(w)
	}
	m, n := len(board), len(board[0])
	out := make([]string, 0, len(words))
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backtrack(board, t, m, n, i, j, &out)
		}
	}
	return out
}

func backtrack(board [][]byte, t *Trie, m, n, x, y int, out *[]string) {
	if x < 0 || m-1 < x || y < 0 || n-1 < y || board[x][y] == '#' {
		return
	}

	if t = t.next[board[x][y]-'a']; t == nil {
		return
	}

	tmp := board[x][y]
	board[x][y] = '#'
	defer func() { board[x][y] = tmp }()

	if t.hasW {
		*out = append(*out, t.word)
		t.hasW = false
	}

	backtrack(board, t, m, n, x-1, y, out)
	backtrack(board, t, m, n, x+1, y, out)
	backtrack(board, t, m, n, x, y-1, out)
	backtrack(board, t, m, n, x, y+1, out)
}

type Trie struct {
	next [26]*Trie
	hasW bool
	word string
}

func (t *Trie) Insert(word string) {
	for i := range word {
		id := word[i] - 'a'
		if t.next[id] == nil {
			t.next[id] = &Trie{}
		}
		t = t.next[id]
	}
	t.hasW = true
	t.word = word
}
