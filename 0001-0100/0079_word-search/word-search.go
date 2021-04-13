package main

/*
  给定一个二维网格和一个单词，找出该单词是否存在于网格中。
  单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

  示例:
    board =
    [
      ['A','B','C','E'],
      ['S','F','C','S'],
      ['A','D','E','E']
    ]

    给定 word = "ABCCED", 返回 true
    给定 word = "SEE", 返回 true
    给定 word = "ABCB", 返回 false

  提示：
    board 和 word 中只包含大写和小写英文字母。
    1 <= board.length <= 200
    1 <= board[i].length <= 200
    1 <= word.length <= 10^3

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/word-search
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking + DFS
func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if backtrack(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, i, row, col int) bool {
	if i == len(word) {
		return true
	}
	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[0])-1 || board[row][col] == '#' || board[row][col] != word[i] {
		return false
	}
	backup := board[row][col]
	board[row][col] = '#'
	if backtrack(board, word, i+1, row-1, col) ||
		backtrack(board, word, i+1, row+1, col) ||
		backtrack(board, word, i+1, row, col-1) ||
		backtrack(board, word, i+1, row, col+1) {
		return true
	}
	board[row][col] = backup
	return false
}
