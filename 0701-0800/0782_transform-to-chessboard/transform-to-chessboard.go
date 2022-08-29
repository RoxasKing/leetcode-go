package main

import "math/bits"

// Difficulty:
// Hard

// Tags:
// Math
// Bit Manipulation

func movesToChessboard(board [][]int) int {
	n := len(board)
	rows, cols := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rows[i] = rows[i]<<1 + board[i][j]
			cols[j] = cols[j]<<1 + board[i][j]
		}
	}
	if bits.OnesCount(uint(rows[0])) != n>>1 && (n&1 == 0 || bits.OnesCount(uint(rows[0])) != n>>1+1) ||
		bits.OnesCount(uint(cols[0])) != n>>1 && (n&1 == 0 || bits.OnesCount(uint(cols[0])) != n>>1+1) {
		return -1
	}
	mask := 1
	for i := 1; i < n; i++ {
		mask = mask<<1 + 1
	}
	rowSwap, colSwap := 0, 0
	for i := 0; i < n; i++ {
		if rows[i] != rows[0] && rows[i] != rows[0]^mask || cols[i] != cols[0] && cols[i] != cols[0]^mask {
			return -1
		}
		if board[i][0] == i&1 {
			rowSwap++
		}
		if board[0][i] == i&1 {
			colSwap++
		}
	}
	if n&1 == 1 {
		if rowSwap&1 == 1 {
			rowSwap = n - rowSwap
		}
		if colSwap&1 == 1 {
			colSwap = n - colSwap
		}
	} else {
		rowSwap = Min(rowSwap, n-rowSwap)
		colSwap = Min(colSwap, n-colSwap)
	}
	return (rowSwap + colSwap) >> 1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
