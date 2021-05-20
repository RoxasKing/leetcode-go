package main

import "strings"

/*
  The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

  Given an integer n, return the number of distinct solutions to the n-queens puzzle.

  Example 1:
    Input: n = 4
    Output: 2
    Explanation: There are two distinct solutions to the 4-queens puzzle as shown.

  Example 2:
    Input: n = 1
    Output: 1

  Constraints:
    1 <= n <= 9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/n-queens-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
// DFS

func totalNQueens(n int) int {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	c := make([]bool, n)
	x := make([]bool, n<<1-1)
	y := make([]bool, n<<1-1)
	out := 0
	dfs(board, c, x, y, 0, n, &out)
	return out
}

func dfs(board []string, c, x, y []bool, i, n int, out *int) {
	if i == n {
		*out++
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
