package main

/*
  On an 8 x 8 chessboard, there is exactly one white rook 'R' and some number of white bishops 'B', black pawns 'p', and empty squares '.'.

  When the rook moves, it chooses one of four cardinal directions (north, east, south, or west), then moves in that direction until it chooses to stop, reaches the edge of the board, captures a black pawn, or is blocked by a white bishop. A rook is considered attacking a pawn if the rook can capture the pawn on the rook's turn. The number of available captures for the white rook is the number of pawns that the rook is attacking.

  Return the number of available captures for the white rook.

  Example 1:
    Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","R",".",".",".","p"],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
    Output: 3
    Explanation: In this example, the rook is attacking all the pawns.

  Example 2:
    Input: board = [[".",".",".",".",".",".",".","."],[".","p","p","p","p","p",".","."],[".","p","p","B","p","p",".","."],[".","p","B","R","B","p",".","."],[".","p","p","B","p","p",".","."],[".","p","p","p","p","p",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
    Output: 0
    Explanation: The bishops are blocking the rook from attacking any of the pawns.

  Example 3:
    Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","p",".",".",".","."],["p","p",".","R",".","p","B","."],[".",".",".",".",".",".",".","."],[".",".",".","B",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."]]
    Output: 3
    Explanation: The rook is attacking the pawns at positions b5, d6, and f5.

  Constraints:
    1. board.length == 8
    2. board[i].length == 8
    3. board[i][j] is either 'R', '.', 'B', or 'p'
    4. There is exactly one cell with board[i][j] == 'R'

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/available-captures-for-rook
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numRookCaptures(board [][]byte) int {
	m, n := len(board), len(board[0])
	x, y := findRook(board)
	out := 0
	for i := x + 1; i < m; i++ {
		if board[i][y] == 'B' {
			break
		} else if board[i][y] == 'p' {
			out++
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		if board[i][y] == 'B' {
			break
		} else if board[i][y] == 'p' {
			out++
			break
		}
	}
	for j := y + 1; j < n; j++ {
		if board[x][j] == 'B' {
			break
		} else if board[x][j] == 'p' {
			out++
			break
		}
	}
	for j := y - 1; j >= 0; j-- {
		if board[x][j] == 'B' {
			break
		} else if board[x][j] == 'p' {
			out++
			break
		}
	}
	return out
}

func findRook(board [][]byte) (int, int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'R' {
				return i, j
			}
		}
	}
	return -1, -1
}
