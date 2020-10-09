package main

/*
  给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

  单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

  说明:
  你可以假设所有输入都由小写字母 a-z 组成。

  提示:
    你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
    如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/word-search-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
