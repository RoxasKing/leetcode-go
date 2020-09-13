package main

/*
  给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
  找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
  示例:
    X X X X
    X O O X
    X X O X
    X O X X
  运行你的函数后，矩阵变为：
    X X X X
    X X X X
    X X X X
    X O X X
  解释:
  被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
*/

// BFS
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	var q [][]int

	checkIfAppend := func(row, col int) {
		if 0 <= row && row < len(board) &&
			0 <= col && col < len(board[0]) &&
			board[row][col] == 'O' {
			q = append(q, []int{row, col})
			board[row][col] = '#'
		}
	}

	isEdge := func(row, col int) bool {
		return row == 0 || row == len(board)-1 || col == 0 || col == len(board[0])-1
	}

	for i := range board {
		for j := range board[0] {
			if isEdge(i, j) && board[i][j] == 'O' {
				q = append(q, []int{i, j})
				board[i][j] = '#'
			}
		}
	}

	for len(q) != 0 {
		row, col := q[0][0], q[0][1]
		q = q[1:]
		checkIfAppend(row-1, col)
		checkIfAppend(row+1, col)
		checkIfAppend(row, col-1)
		checkIfAppend(row, col+1)
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// DFS + Recursion
func solve2(board [][]byte) {
	if len(board) == 0 {
		return
	}
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= len(board) ||
			col < 0 || col >= len(board[0]) ||
			board[row][col] == 'X' || board[row][col] == '#' {
			return
		}
		board[row][col] = '#'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}
	isEdge := func(row, col int) bool {
		return row == 0 || row == len(board)-1 || col == 0 || col == len(board[0])-1
	}
	for i := range board {
		for j := range board[0] {
			if isEdge(i, j) && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// DFS + Stack
func solve3(board [][]byte) {
	if len(board) == 0 {
		return
	}
	type coordinate struct {
		row int
		col int
	}
	dfs := func(row, col int) {
		var stack []coordinate
		stack = append(stack, coordinate{row, col})
		board[row][col] = '#'
		for len(stack) != 0 {
			cur := stack[len(stack)-1]
			if cur.row-1 >= 0 && board[cur.row-1][cur.col] == 'O' {
				stack = append(stack, coordinate{cur.row - 1, cur.col})
				board[cur.row-1][cur.col] = '#'
				continue
			}
			if cur.row+1 <= len(board)-1 && board[cur.row+1][cur.col] == 'O' {
				stack = append(stack, coordinate{cur.row + 1, cur.col})
				board[cur.row+1][cur.col] = '#'
				continue
			}
			if cur.col-1 >= 0 && board[cur.row][cur.col-1] == 'O' {
				stack = append(stack, coordinate{cur.row, cur.col - 1})
				board[cur.row][cur.col-1] = '#'
				continue
			}
			if cur.col+1 <= len(board[0])-1 && board[cur.row][cur.col+1] == 'O' {
				stack = append(stack, coordinate{cur.row, cur.col + 1})
				board[cur.row][cur.col+1] = '#'
				continue
			}
			stack = stack[:len(stack)-1]
		}
	}
	for i := range board {
		for j := range board[0] {
			isEdge := i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1
			if isEdge && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// union-find
func solve4(board [][]byte) {
	if len(board) == 0 {
		return
	}
	dummyNode := len(board) * len(board[0])
	parents := make([]int, len(board)*len(board[0])+1)
	for i := range parents {
		parents[i] = i
	}
	find := func(node int) int {
		for parents[node] != node {
			parents[node] = parents[parents[node]]
			node = parents[node]
		}
		return node
	}
	union := func(node1, node2 int) {
		root1 := find(node1)
		root2 := find(node2)
		if root1 != root2 {
			parents[root2] = root1
		}
	}
	isConnected := func(node1, node2 int) bool {
		return find(node1) == find(node2)
	}
	node := func(row, col int) int {
		return row*len(board[0]) + col
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1 {
					union(node(i, j), dummyNode)
				} else {
					if i > 0 && board[i-1][j] == 'O' {
						union(node(i, j), node(i-1, j))
					}
					if i < len(board)-1 && board[i+1][j] == 'O' {
						union(node(i, j), node(i+1, j))
					}
					if j > 0 && board[i][j-1] == 'O' {
						union(node(i, j), node(i, j-1))
					}
					if j < len(board[0])-1 && board[i][j+1] == 'O' {
						union(node(i, j), node(i, j+1))
					}
				}
			}
		}
	}
	for i := range board {
		for j := range board[0] {
			if isConnected(node(i, j), dummyNode) {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
