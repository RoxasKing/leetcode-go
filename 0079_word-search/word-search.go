package main

/*
  给定一个二维网格和一个单词，找出该单词是否存在于网格中。

  单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

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
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var backtrack func(int, int, int) bool
	backtrack = func(row, col, index int) bool {
		if index == len(word) {
			return true
		}
		backup := board[row][col]
		board[row][col] = '#'
		for _, move := range moves {
			newRow, newCol := row+move[0], col+move[1]
			if newRow < 0 || len(board)-1 < newRow ||
				newCol < 0 || len(board[0])-1 < newCol ||
				board[newRow][newCol] != word[index] {
				continue
			}
			if backtrack(newRow, newCol, index+1) {
				return true
			}
		}
		board[row][col] = backup
		return false
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] && backtrack(i, j, 1) {
				return true
			}
		}
	}
	return false
}
