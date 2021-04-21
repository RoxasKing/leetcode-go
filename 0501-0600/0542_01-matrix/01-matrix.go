package main

/*
  Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.

  The distance between two adjacent cells is 1.

  Example 1:
    Input:
    [[0,0,0],
     [0,1,0],
     [0,0,0]]
    Output:
    [[0,0,0],
     [0,1,0],
     [0,0,0]]

  Example 2:
    Input:
    [[0,0,0],
     [0,1,0],
     [1,1,1]]
    Output:
    [[0,0,0],
     [0,1,0],
     [1,2,1]]

  Note:
    1. The number of elements of the given matrix will not exceed 10,000.
    2. There are at least one 0 in the given matrix.
    3. The cells are adjacent in only four directions: up, down, left and right.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/01-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS
func updateMatrix(matrix [][]int) [][]int {
	forwards := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(matrix), len(matrix[0])
	out := make([][]int, m)
	for i := range out {
		out[i] = make([]int, n)
	}

	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				q = append(q, []int{i, j})
			}
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		i, j := e[0], e[1]
		for _, f := range forwards {
			x, y := i+f[0], j+f[1]
			if x < 0 || x > m-1 || y < 0 || y > n-1 || matrix[x][y] == 0 || out[x][y] != 0 {
				continue
			}
			out[x][y] = out[i][j] + 1
			q = append(q, []int{x, y})
		}
	}

	return out
}

// Dynamic Programming
func updateMatrix2(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	out := make([][]int, m)
	for i := range out {
		out[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				out[i][j] = 0
			} else {
				out[i][j] = 1<<31 - 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				out[i][j] = Min(out[i][j], out[i-1][j]+1)
			}
			if j > 0 {
				out[i][j] = Min(out[i][j], out[i][j-1]+1)
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 {
				out[i][j] = Min(out[i][j], out[i+1][j]+1)
			}
			if j < n-1 {
				out[i][j] = Min(out[i][j], out[i][j+1]+1)
			}
		}
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
