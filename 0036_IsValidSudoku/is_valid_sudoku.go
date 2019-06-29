package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	dict := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] >= '1' && board[i][j] <= '9' {
				if _, ok := dict[board[i][j]]; !ok {
					dict[board[i][j]] = true
				} else {
					return false
				}
			}
		}
		dict = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] >= '1' && board[j][i] <= '9' {
				if _, ok := dict[board[j][i]]; !ok {
					dict[board[j][i]] = true
				} else {
					return false
				}
			}
		}
		dict = make(map[byte]bool)
	}
	for i := 0; i <= 6; i += 3 {
		for j := 0; j <= 6; j += 3 {
			cube := []byte{
				board[i][j],
				board[i][j+1],
				board[i][j+2],
				board[i+1][j],
				board[i+1][j+1],
				board[i+1][j+2],
				board[i+2][j],
				board[i+2][j+1],
				board[i+2][j+2],
			}
			for _, elem := range cube {
				if elem >= '1' && elem <= '9' {
					if _, ok := dict[elem]; !ok {
						dict[elem] = true
					} else {
						return false
					}
				}
			}
			dict = make(map[byte]bool)
		}
	}
	return true
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
	fmt.Println(isValidSudoku(in))
}
