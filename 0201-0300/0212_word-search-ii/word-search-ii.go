package main

// Difficulty:
// Hard

// Tags:
// Trie
// Backtracking

func findWords(board [][]byte, words []string) []string {
	t := &Trie{}
	for _, w := range words {
		t.Insert(w)
	}
	out := []string{}
	for i := range board {
		for j := range board[i] {
			dfs(board, t, &out, i, j)
		}
	}
	return out
}

func dfs(b [][]byte, t *Trie, out *[]string, i, j int) {
	if i < 0 || len(b)-1 < i || j < 0 || len(b[0])-1 < j || b[i][j] == '#' || t.next[b[i][j]-'a'] == nil {
		return
	}
	t = t.next[b[i][j]-'a']
	if t.hasW {
		*out = append(*out, t.word)
		t.hasW = false
	}
	tmp := b[i][j]
	b[i][j] = '#'
	defer func() { b[i][j] = tmp }()
	dfs(b, t, out, i-1, j)
	dfs(b, t, out, i+1, j)
	dfs(b, t, out, i, j-1)
	dfs(b, t, out, i, j+1)
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
