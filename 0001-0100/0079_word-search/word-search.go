package main

/*
  Given an m x n grid of characters board and a string word, return true if word exists in the grid.

  The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

  Example 1:
    Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
    Output: true

  Example 2:
    Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
    Output: true

  Example 3:
    Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
    Output: false

  Constraints:
    1. m == board.length
    2. n = board[i].length
    3. 1 <= m, n <= 6
    4. 1 <= word.length <= 15
    5. board and word consists of only lowercase and uppercase English letters.

  Follow up: Could you use search pruning to make your solution faster with a larger board?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/word-search
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, m, n, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, m, n int, word string, i, x, y int) bool {
	if i == len(word) {
		return true
	}

	if x < 0 || x > m-1 || y < 0 || y > n-1 || board[x][y] != word[i] {
		return false
	}

	ch := board[x][y]
	board[x][y] = '$'

	if dfs(board, m, n, word, i+1, x-1, y) ||
		dfs(board, m, n, word, i+1, x+1, y) ||
		dfs(board, m, n, word, i+1, x, y-1) ||
		dfs(board, m, n, word, i+1, x, y+1) {
		return true
	}

	board[x][y] = ch

	return false
}
