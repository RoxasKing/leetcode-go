package main

/*
  编写一个程序，通过填充空格来解决数独问题。
  一个数独的解法需遵循如下规则：
    数字 1-9 在每一行只能出现一次。
    数字 1-9 在每一列只能出现一次。
    数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
  空白格用 '.' 表示。

  一个数独。
  答案被标成红色。

  提示：
    给定的数独序列只包含数字 1-9 和字符 '.' 。
    你可以假设给定的数独只有唯一解。
    给定数独永远是 9x9 形式的。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sudoku-solver
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func solveSudoku(board [][]byte) {
	row := [9][9]int{}
	col := [9][9]int{}
	box := [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				k := (i/3)*3 + j/3
				num := board[i][j] - '1'
				row[i][num] = 1
				col[j][num] = 1
				box[k][num] = 1
			}
		}
	}
	_ = backtrack(board, row, col, box, 0)
}

func backtrack(board [][]byte, row, col, box [9][9]int, index int) bool {
	if index == 81 {
		return true
	}
	i, j := index/9, index%9
	if board[i][j] != '.' {
		return backtrack(board, row, col, box, index+1)
	}
	k := (i/3)*3 + j/3
	for num := 0; num < 9; num++ {
		if row[i][num]+col[j][num]+box[k][num] == 0 {
			row[i][num] = 1
			col[j][num] = 1
			box[k][num] = 1

			board[i][j] = byte(num) + '1'
			if backtrack(board, row, col, box, index+1) {
				return true
			}
			board[i][j] = '.'

			row[i][num] = 0
			col[j][num] = 0
			box[k][num] = 0
		}
	}
	return false
}
