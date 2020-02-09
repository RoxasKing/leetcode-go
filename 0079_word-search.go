package leetcode

/*
  给定一个二维网格和一个单词，找出该单词是否存在于网格中。
  单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用
*/

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	var dfs func(int, int, int) bool
	dfs = func(row, col, index int) bool {
		if len(word) == index {
			return true
		}
		if row < 0 || row >= len(board) ||
			col < 0 || col >= len(board[0]) ||
			board[row][col] != word[index] {
			return false
		}
		tmp := board[row][col]
		board[row][col] = 0
		if dfs(row-1, col, index+1) ||
			dfs(row+1, col, index+1) ||
			dfs(row, col-1, index+1) ||
			dfs(row, col+1, index+1) {
			return true
		}
		board[row][col] = tmp
		return false
	}
	for i := range board {
		for j := range board[0] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
