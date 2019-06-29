package main

import "fmt"

func solveSudoku(board [][]byte) {
	rows := make(map[int]map[byte]bool)
	cows := make(map[int]map[byte]bool)
	//cubes := make(map[int]map[byte]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] >= '1' && board[i][j] <= '9' {
				row := rows[i]
				if row == nil {
					row = make(map[byte]bool)
					row[board[i][j]] = true
					rows[i] = row
				}
			}
			if board[j][i] >= '1' && board[j][i] <= '9' {
				cow := cows[j]
				if cow == nil {
					cow = make(map[byte]bool)
					cow[board[j][i]] = true
					cows[j] = cow
				}
			}
		}
	}
	fmt.Println(rows)
	fmt.Println(cows)
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
	solveSudoku(in)
}
