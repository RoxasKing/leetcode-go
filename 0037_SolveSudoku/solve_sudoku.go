package main

import "fmt"

func solveSudoku(board [][]byte) {
	solve(board, 0)
}

func solve(board [][]byte, k int) bool {
	// k 是把 9x9 矩阵转为一维数组后的值
	if k == 81 {
		return true
	}
	// 将 k 转换为二维数组的坐标
	r, c := k/9, k%9
	if board[r][c] != '.' {
		// 如果当前坐标非空白，则继续解下一个
		return solve(board, k+1)
	}
	// bi,bj 是 r,c所在 3x3 矩阵中的左上角坐标
	bi, bj := r/3*3, c/3*3
	// 匿名函数，用于检查 b 是否能放 r,c 位
	isValid := func(b byte) bool {
		for n := 0; n < 9; n++ {
			if board[r][n] == b || board[n][c] == b || board[bi+n/3][bj+n%3] == b {
				return false
			}
		}
		return true
	}
	// 分别填入 ‘1’ ~ ‘9’ ，检查是否合规
	for b := byte('1'); b <= '9'; b++ {
		if isValid(b) {
			// 如果当前位置允许
			board[r][c] = b
			// 则检查填入当前数字后，下一个位置是否也可以填入合规数字
			if solve(board, k+1) {
				// 只有下一个位置依然能填入合规数字的情况下，当前位置才能返回 true
				return true
			}
		}
	}
	// 不合规则擦除,并返回 false
	board[r][c] = '.'
	return false
}

func main() {
	in := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(in)
	solveSudoku(in)
	fmt.Println(in)
}
