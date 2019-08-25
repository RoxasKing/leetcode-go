package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	if len(word) == 0 {
		return false
	}

	var dfs func(int, int, int) bool
	dfs = func(r int, c int, index int) bool {
		if len(word) == index {
			return true
		}
		// board[r][c] != word[index] 和下面 boar[r][c]=0结合使用
		if r < 0 || m <= r || c < 0 || n <= c || board[r][c] != word[index] {
			return false
		}
		temp := board[r][c]
		// 将经过的路径暂时标记为 0 ，防止重复经过
		board[r][c] = 0

		if dfs(r-1, c, index+1) ||
			dfs(r+1, c, index+1) ||
			dfs(r, c-1, index+1) ||
			dfs(r, c+1, index+1) {
			return true
		}
		board[r][c] = temp
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCBA"
	fmt.Println(exist(board, word))
}
