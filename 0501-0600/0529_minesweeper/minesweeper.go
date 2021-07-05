package main

// Tags:
// DFS + Recursion
func updateBoard(board [][]byte, click []int) [][]byte {
	out := make([][]byte, len(board))
	for i := range out {
		out[i] = make([]byte, len(board[0]))
		copy(out[i], board[i])
	}
	if out[click[0]][click[1]] == 'M' {
		out[click[0]][click[1]] = 'X'
	} else {
		out[click[0]][click[1]] = 'B'
		recur(out, click[0], click[1])
	}
	return out
}

var locals = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func recur(board [][]byte, row, col int) {
	var count int
	var lists [][]int
	for _, local := range locals {
		newRow, newCol := row+local[0], col+local[1]
		if 0 <= newRow && newRow < len(board) && 0 <= newCol && newCol < len(board[0]) {
			if board[newRow][newCol] == 'M' {
				count++
			} else if board[newRow][newCol] == 'E' {
				lists = append(lists, []int{newRow, newCol})
			}
		}
	}
	if count != 0 {
		board[row][col] = byte(count) + '0'
		return
	}
	for _, e := range lists {
		board[e[0]][e[1]] = 'B'
		recur(board, e[0], e[1])
	}
}

// BFS + Interation
func updateBoard2(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	board[click[0]][click[1]] = 'B'
	queue := [][]int{click}
	for len(queue) != 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]
		var count int
		var clicks [][]int
		for _, local := range locals {
			newRow, newCol := row+local[0], col+local[1]
			if 0 <= newRow && newRow < len(board) && 0 <= newCol && newCol < len(board[0]) {
				if board[newRow][newCol] == 'M' {
					count++
				} else if board[newRow][newCol] == 'E' {
					clicks = append(clicks, []int{newRow, newCol})
				}
			}
		}
		if count != 0 {
			board[row][col] = byte(count) + '0'
			continue
		}
		for _, click := range clicks {
			board[click[0]][click[1]] = 'B'
		}
		queue = append(queue, clicks...)
	}
	return board
}
