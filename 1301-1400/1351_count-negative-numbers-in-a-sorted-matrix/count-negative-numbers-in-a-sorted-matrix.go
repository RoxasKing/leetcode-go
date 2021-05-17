package main

/*
  Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.

  Example 1:
    Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
    Output: 8
    Explanation: There are 8 negatives number in the matrix.

  Example 2:
    Input: grid = [[3,2],[1,0]]
    Output: 0

  Example 3:
    Input: grid = [[1,-1],[-1,-1]]
    Output: 3

  Example 4:
    Input: grid = [[-1]]
    Output: 1

  Constraints:
    1. m == grid.length
    2. n == grid[i].length
    3. 1 <= m, n <= 100
    4. -100 <= grid[i][j] <= 100

  Follow up: Could you find an O(n + m) solution?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers

func countNegatives(grid [][]int) int {
	out := 0
	m, n := len(grid), len(grid[0])
	for i, j := m-1, 0; i >= 0 && j < n; {
		if grid[i][j] < 0 {
			out += n - j
			i--
		} else {
			j++
		}
	}
	return out
}
