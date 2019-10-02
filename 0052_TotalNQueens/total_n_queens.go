package main

import "fmt"

var count int

func totalNQueens(n int) int {
	count = 0
	if n < 1 {
		return count
	}
	DFS(n, 0, 0, 0, 0)
	return count
}

func DFS(n, row, col, pie, na int) {
	if row == n {
		count += 1
		return
	}
	bits := (^(col | pie | na)) & ((1 << uint(n)) - 1)
	for bits > 0 {
		p := bits & -bits
		DFS(n, row+1, col|p, (pie|p)<<1, (na|p)>>1)
		bits &= bits - 1
	}
}

func main() {
	fmt.Println(totalNQueens(4))
}
