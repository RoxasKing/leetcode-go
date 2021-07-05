package main

// Tags:
// Backtracking + Trie
func findWords(board [][]byte, words []string) []string {
	trie := buildTrie(words)
	contain := make(map[string]bool)
	out := []string{}
	for i := range board {
		for j := range board[0] {
			backtrack(board, trie, contain, "", &out, i, j)
		}
	}
	return out
}

func backtrack(board [][]byte, t *Trie, contain map[string]bool, cur string, out *[]string, row, col int) {
	if row < 0 || len(board)-1 < row || col < 0 || len(board[0])-1 < col || board[row][col] == '#' ||
		t.next[board[row][col]-'a'] == nil {
		return
	}
	cur += string(board[row][col])
	t = t.next[board[row][col]-'a']

	save := board[row][col]
	board[row][col] = '#'

	if t.isEnd && !contain[cur] {
		*out = append(*out, cur)
		contain[cur] = true
	}

	backtrack(board, t, contain, cur, out, row-1, col)
	backtrack(board, t, contain, cur, out, row+1, col)
	backtrack(board, t, contain, cur, out, row, col-1)
	backtrack(board, t, contain, cur, out, row, col+1)

	board[row][col] = save
}

func buildTrie(words []string) *Trie {
	trie := new(Trie)
	for _, word := range words {
		t := trie
		for i := range word {
			index := word[i] - 'a'
			if t.next[index] == nil {
				t.next[index] = new(Trie)
			}
			t = t.next[index]
		}
		t.isEnd = true
	}
	return trie
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}
