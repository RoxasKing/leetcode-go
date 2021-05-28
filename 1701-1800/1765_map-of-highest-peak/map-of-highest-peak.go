package main

/*
  You are given an integer matrix isWater of size m x n that represents a map of land and water cells.

    1. If isWater[i][j] == 0, cell (i, j) is a land cell.
    2. If isWater[i][j] == 1, cell (i, j) is a water cell.

  You must assign each cell a height in a way that follows these rules:

    1. The height of each cell must be non-negative.
    2. If the cell is a water cell, its height must be 0.
    3. Any two adjacent cells must have an absolute height difference of at most 1. A cell is adjacent to another cell if the former is directly north, east, south, or west of the latter (i.e., their sides are touching).

  Find an assignment of heights such that the maximum height in the matrix is maximized.

  Return an integer matrix height of size m x n where height[i][j] is cell (i, j)'s height. If there are multiple solutions, return any of them.

  Example 1:
    Input: isWater = [[0,1],[0,0]]
    Output: [[1,0],[2,1]]
    Explanation: The image shows the assigned heights of each cell.
      The blue cell is the water cell, and the green cells are the land cells.

  Example 2:
    Input: isWater = [[0,0,1],[1,0,0],[0,0,0]]
    Output: [[1,1,0],[0,1,1],[1,2,2]]
    Explanation: A height of 2 is the maximum possible height of any assignment.
      Any height assignment that has a maximum height of 2 while still meeting the rules will also be accepted.

  Constraints:
    1. m == isWater.length
    2. n == isWater[i].length
    3. 1 <= m, n <= 1000
    4. isWater[i][j] is 0 or 1.
    5. There is at least one water cell.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/map-of-highest-peak
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	out := make([][]int, m)
	visited := make([][]bool, m)
	for i := range out {
		out[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	q := [][3]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				q = append(q, [3]int{i, j, 0})
			}
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		i, j, h := e[0], e[1], e[2]
		for _, f := range forwards {
			ni, nj := i+f[0], j+f[1]
			if ni < 0 || m-1 < ni || nj < 0 || n-1 < nj || isWater[ni][nj] == 1 || visited[ni][nj] {
				continue
			}
			visited[ni][nj] = true
			out[ni][nj] = h + 1
			q = append(q, [3]int{ni, nj, h + 1})
		}
	}

	return out
}

var forwards = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
