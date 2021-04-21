package main

/*
  You are given an n x n grid where you have placed some 1 x 1 x 1 cubes. Each value v = grid[i][j] represents a tower of v cubes placed on top of cell (i, j).

  After placing these cubes, you have decided to glue any directly adjacent cubes to each other, forming several irregular 3D shapes.

  Return the total surface area of the resulting shapes.

  Note: The bottom face of each shape counts toward its surface area.

  Example 1:
    Input: grid = [[2]]
    Output: 10

  Example 2:
    Input: grid = [[1,2],[3,4]]
    Output: 34

  Example 3:
    Input: grid = [[1,0],[0,2]]
    Output: 16

  Example 4:
    Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
    Output: 32

  Example 5:
    Input: grid = [[2,2,2],[2,1,2],[2,2,2]]
    Output: 46

  Constraints:
    1. n == grid.length
    2. n == grid[i].length
    3. 1 <= n <= 50
    4. 0 <= grid[i][j] <= 50

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/surface-area-of-3d-shapes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func surfaceArea(grid [][]int) int {
	n := len(grid)
	out := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			out += 2
			if i == 0 {
				out += grid[i][j]
			}
			if i == n-1 {
				out += grid[i][j]
			}
			if j == 0 {
				out += grid[i][j]
			}
			if j == n-1 {
				out += grid[i][j]
			}
			if i > 0 && grid[i][j] > grid[i-1][j] {
				out += grid[i][j] - grid[i-1][j]
			}
			if i < n-1 && grid[i][j] > grid[i+1][j] {
				out += grid[i][j] - grid[i+1][j]
			}
			if j > 0 && grid[i][j] > grid[i][j-1] {
				out += grid[i][j] - grid[i][j-1]
			}
			if j < n-1 && grid[i][j] > grid[i][j+1] {
				out += grid[i][j] - grid[i][j+1]
			}
		}
	}
	return out
}
