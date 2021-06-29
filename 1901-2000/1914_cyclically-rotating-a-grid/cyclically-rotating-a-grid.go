package main

/*
  You are given an m x n integer matrix grid​​​, where m and n are both even integers, and an integer k.

  The matrix is composed of several layers, which is shown in the below image, where each color is its own layer:

  A cyclic rotation of the matrix is done by cyclically rotating each layer in the matrix. To cyclically rotate a layer once, each element in the layer will take the place of the adjacent element in the counter-clockwise direction. An example rotation is shown below:

  Return the matrix after applying k cyclic rotations to it.

  Example 1:
    Input: grid = [[40,10],[30,20]], k = 1
    Output: [[10,20],[40,30]]
    Explanation: The figures above represent the grid at every state.

  Example 2:
    Input: grid = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]], k = 2
    Output: [[3,4,8,12],[2,11,10,16],[1,7,6,15],[5,9,13,14]]
    Explanation: The figures above represent the grid at every state.

  Constraints:
    1. m == grid.length
    2. n == grid[i].length
    3. 2 <= m, n <= 50
    4. Both m and n are even integers.
    5. 1 <= grid[i][j] <= 5000
    6. 1 <= k <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/cyclically-rotating-a-grid
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	arr := [][]int{}
	for l, r, u, d := 0, n-1, 0, m-1; l <= r && u <= d; {
		cur := []int{}
		for i := l; i <= r; i++ {
			cur = append(cur, grid[u][i])
		}
		u++
		for i := u; i <= d; i++ {
			cur = append(cur, grid[i][r])
		}
		r--
		for i := r; i >= l; i-- {
			cur = append(cur, grid[d][i])
		}
		d--
		for i := d; i >= u; i-- {
			cur = append(cur, grid[i][l])
		}
		l++
		arr = append(arr, cur)
	}
	for i := range arr {
		x := k % len(arr[i])
		arr[i] = append(arr[i][x:], arr[i][:x]...)
	}
	for i, l, r, u, d := 0, 0, n-1, 0, m-1; l <= r && u <= d; i++ {
		cur := arr[i]
		for i := l; i <= r; i++ {
			grid[u][i] = cur[0]
			cur = cur[1:]
		}
		u++
		for i := u; i <= d; i++ {
			grid[i][r] = cur[0]
			cur = cur[1:]
		}
		r--
		for i := r; i >= l; i-- {
			grid[d][i] = cur[0]
			cur = cur[1:]
		}
		d--
		for i := d; i >= u; i-- {
			grid[i][l] = cur[0]
			cur = cur[1:]
		}
		l++
	}
	return grid
}
