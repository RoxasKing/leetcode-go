package main

import (
	"container/heap"
)

/*
  A rhombus sum is the sum of the elements that form the border of a regular rhombus shape in grid​​​. The rhombus must have the shape of a square rotated 45 degrees with each of the corners centered in a grid cell. Below is an image of four valid rhombus shapes with the corresponding colored cells that should be included in each rhombus sum:

  Note that the rhombus can have an area of 0, which is depicted by the purple rhombus in the bottom right corner.

  Return the biggest three distinct rhombus sums in the grid in descending order. If there are less than three distinct values, return all of them.

  Example 1:
    Input: grid = [[3,4,5,1,3],[3,3,4,2,3],[20,30,200,40,10],[1,5,5,4,1],[4,3,2,2,5]]
    Output: [228,216,211]
    Explanation: The rhombus shapes for the three biggest distinct rhombus sums are depicted above.
      - Blue: 20 + 3 + 200 + 5 = 228
      - Red: 200 + 2 + 10 + 4 = 216
      - Green: 5 + 200 + 4 + 2 = 211

  Example 2:
    Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
    Output: [20,9,8]
    Explanation: The rhombus shapes for the three biggest distinct rhombus sums are depicted above.
      - Blue: 4 + 2 + 6 + 8 = 20
      - Red: 9 (area 0 rhombus in the bottom right corner)
      - Green: 8 (area 0 rhombus in the bottom middle)

  Example 3:
    Input: grid = [[7,7,7]]
    Output: [7]
    Explanation: All three possible rhombus sums are the same, so return [7].

  Constraints:
    1. m == grid.length
    2. n == grid[i].length
    3. 1 <= m, n <= 50
    4. 1 <= grid[i][j] <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/get-biggest-three-rhombus-sums-in-a-grid
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue

func getBiggestThree(grid [][]int) []int {
	h := MaxHeap{}
	m, n := len(grid), len(grid[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			heap.Push(&h, grid[r][c])
			for i := 1; c-i >= 0 && c+i < n && r+2*i < m; i++ {
				sum := grid[r][c] + grid[r+i][c-i] + grid[r+i][c+i] + grid[r+2*i][c]
				for j := 1; j < i; j++ {
					sum += grid[r+j][c-j] + grid[r+j][c+j] + grid[r+2*i-j][c-j] + grid[r+2*i-j][c+j]
				}
				heap.Push(&h, sum)
			}
		}
	}
	set := map[int]bool{}
	out := make([]int, 0, 3)
	for len(out) < 3 && h.Len() > 0 {
		num := heap.Pop(&h).(int)
		if !set[num] {
			out = append(out, num)
			set[num] = true
		}
	}
	return out
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
