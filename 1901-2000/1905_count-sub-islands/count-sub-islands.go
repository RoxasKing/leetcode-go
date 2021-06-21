package main

/*
  You are given two m x n binary matrices grid1 and grid2 containing only 0's (representing water) and 1's (representing land). An island is a group of 1's connected 4-directionally (horizontal or vertical). Any cells outside of the grid are considered water cells.

  An island in grid2 is considered a sub-island if there is an island in grid1 that contains all the cells that make up this island in grid2.

  Return the number of islands in grid2 that are considered sub-islands.

  Example 1:
    Input: grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
    Output: 3
    Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
      The 1s colored red in grid2 are those considered to be part of a sub-island. There are three sub-islands.

  Example 2:
    Input: grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
    Output: 2
    Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
      The 1s colored red in grid2 are those considered to be part of a sub-island. There are two sub-islands.

  Constraints:
    1. m == grid1.length == grid2.length
    2. n == grid1[i].length == grid2[i].length
    3. 1 <= m, n <= 500
    4. grid1[i][j] and grid2[i][j] are either 0 or 1.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-sub-islands
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// DFS

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 0 || grid1[i][j] == 0 {
				continue
			}
			grid2[i][j] = 0
			if isSubIsland(grid1, grid2, m, n, i, j) {
				out++
			}
		}
	}
	return out
}

func isSubIsland(grid1, grid2 [][]int, m, n, i, j int) bool {
	out := true
	for _, e := range moves {
		x, y := i+e[0], j+e[1]
		if x < 0 || x > m-1 || y < 0 || y > n-1 || grid2[x][y] == 0 {
			continue
		}
		grid2[x][y] = 0
		out = grid1[x][y] == 1 && out
		out = isSubIsland(grid1, grid2, m, n, x, y) && out
	}
	return out
}

var moves = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
