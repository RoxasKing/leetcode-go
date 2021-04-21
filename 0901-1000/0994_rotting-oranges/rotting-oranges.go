package main

/*
  You are given an m x n grid where each cell can have one of three values:
    1. 0 representing an empty cell,
    2. 1 representing a fresh orange, or
    3. 2 representing a rotten orange.
  Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

  Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

  Example 1:
    Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
    Output: 4

  Example 2:
    Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
    Output: -1
    Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

  Example 3:
    Input: grid = [[0,2]]
    Output: 0
    Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.

  Constraints:
    1. m == grid.length
    2. n == grid[i].length
    3. 1 <= m, n <= 10
    4. grid[i][j] is 0, 1, or 2.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/rotting-oranges
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS
func orangesRotting(grid [][]int) int {
	forwards := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	q := [][]int{}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	out := 0
	for len(q) > 0 {
		size := len(q)
		for _, e := range q {
			x, y := e[0], e[1]
			for _, f := range forwards {
				nx, ny := x+f[0], y+f[1]
				if nx < 0 || m-1 < nx || ny < 0 || n-1 < ny || grid[nx][ny] != 1 {
					continue
				}
				grid[nx][ny] = 2
				cnt--
				q = append(q, []int{nx, ny})
			}
		}
		q = q[size:]
		if len(q) > 0 {
			out++
		}
	}
	if cnt > 0 {
		return -1
	}
	return out
}
