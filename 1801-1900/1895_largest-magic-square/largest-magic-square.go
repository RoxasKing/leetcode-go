package main

/*
  A k x k magic square is a k x k grid filled with integers such that every row sum, every column sum, and both diagonal sums are all equal. The integers in the magic square do not have to be distinct. Every 1 x 1 grid is trivially a magic square.

  Given an m x n integer grid, return the size (i.e., the side length k) of the largest magic square that can be found within this grid.

  Example 1:
    Input: grid = [[7,1,4,5,6],[2,5,1,6,4],[1,5,4,3,2],[1,2,7,3,4]]
    Output: 3
    Explanation: The largest magic square has a size of 3.
      Every row sum, column sum, and diagonal sum of this magic square is equal to 12.
      - Row sums: 5+1+6 = 5+4+3 = 2+7+3 = 12
      - Column sums: 5+5+2 = 1+4+7 = 6+3+3 = 12
      - Diagonal sums: 5+4+3 = 6+4+2 = 12

  Example 2:
    Input: grid = [[5,1,3,1],[9,3,3,1],[1,3,3,8]]
    Output: 2

  Constraints:
    1. m == grid.length
    2. n == grid[i].length
    3. 1 <= m, n <= 50
    4. 1 <= grid[i][j] <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/largest-magic-square
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Prefix Sum

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		f[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			f[i+1][j+1] = grid[i][j] + f[i][j+1] + f[i+1][j] - f[i][j]
		}
	}

	out := 1
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			for k := out + 1; x+k <= m && y+k <= n; k++ {
				if valid(grid, f, x, y, k) {
					out = k
				}
			}
		}
	}
	return out
}

func valid(grid, f [][]int, x, y, k int) bool {
	sum := f[x+k][y+k] - f[x+k][y] - f[x][y+k] + f[x][y]
	if sum%k > 0 {
		return false
	}

	target := sum / k
	d1, d2 := 0, 0

	for i := 0; i < k; i++ {
		if f[x+i+1][y+k]-f[x+i+1][y]-f[x+i][y+k]+f[x+i][y] != target {
			return false
		}
		if f[x+k][y+i+1]-f[x+k][y+i]-f[x][y+i+1]+f[x][y+i] != target {
			return false
		}
		d1 += grid[x+i][y+i]
		d2 += grid[x+k-1-i][y+i]
	}

	return d1 == target && d2 == target
}
