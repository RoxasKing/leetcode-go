package main

import "strings"

/*
  The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

  Given an integer n, return all distinct solutions to the n-queens puzzle.

  Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

  Example 1:
    Input: n = 4
    Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
    Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

  Example 2:
    Input: n = 1
    Output: [["Q"]]

  Constraints:
    1 <= n <= 9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/n-queens
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Backtracking
// DFS

func solveNQueens(n int) [][]string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	c := make([]bool, n)
	x := make([]bool, n<<1-1)
	y := make([]bool, n<<1-1)
	out := [][]string{}
	dfs(board, c, x, y, 0, n, &out)
	return out
}

func dfs(board []string, c, x, y []bool, i, n int, out *[][]string) {
	if i == n {
		tmp := make([]string, n)
		copy(tmp, board)
		*out = append(*out, tmp)
		return
	}

	chs := []byte(board[i])
	for j := 0; j < n; j++ {
		if c[j] || x[i+j] || y[i-j+n-1] {
			continue
		}

		chs[j] = 'Q'
		board[i] = string(chs)
		c[j] = true
		x[i+j] = true
		y[i-j+n-1] = true

		dfs(board, c, x, y, i+1, n, out)

		chs[j] = '.'
		board[i] = string(chs)
		c[j] = false
		x[i+j] = false
		y[i-j+n-1] = false
	}
}
