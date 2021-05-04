package main

/*
  Given an m x n matrix board where each cell is a battleship 'X' or empty '.', return the number of the battleships on board.

  Battleships can only be placed horizontally or vertically on board. In other words, they can only be made of the shape 1 x k (1 row, k columns) or k x 1 (k rows, 1 column), where k can be of any size. At least one horizontal or vertical cell separates between two battleships (i.e., there are no adjacent battleships).

  Example 1:
    Input: board = [["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]
    Output: 2

  Example 2:
    Input: board = [["."]]
    Output: 0

  Constraints:
    1. m == board.length
    2. n == board[i].length
    3. 1 <= m, n <= 200
    4. board[i][j] is either '.' or 'X'.

  Follow up: Could you do it in one-pass, using only O(1) extra memory and without modifying the values board?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/battleships-in-a-board
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countBattleships(board [][]byte) int {
	out := 0
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			if (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				out++
			}
		}
	}
	return out
}
