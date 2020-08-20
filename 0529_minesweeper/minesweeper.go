package main

/*
  让我们一起来玩扫雷游戏！

  给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。

  现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：
    如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
    如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的未挖出方块都应该被递归地揭露。
    如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
    如果在此次点击中，若无更多方块可被揭露，则返回面板。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minesweeper
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
