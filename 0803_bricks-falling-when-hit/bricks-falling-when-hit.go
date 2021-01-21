package main

/*
  You are given an m x n binary grid, where each 1 represents a brick and 0 represents an empty space. A brick is stable if:
    It is directly connected to the top of the grid, or
    At least one other brick in its four adjacent cells is stable.
  You are also given an array hits, which is a sequence of erasures we want to apply. Each time we want to erase the brick at the location hits[i] = (rowi, coli). The brick on that location (if it exists) will disappear. Some other bricks may no longer be stable because of that erasure and will fall. Once a brick falls, it is immediately erased from the grid (i.e., it does not land on other stable bricks).
  Return an array result, where each result[i] is the number of bricks that will fall after the ith erasure is applied.

  Note that an erasure may refer to a location with no brick, and if it does, no bricks drop.

  Example 1:
    Input: grid = [[1,0,0,0],[1,1,1,0]], hits = [[1,0]]
    Output: [2]
    Explanation: Starting with the grid:
    [[1,0,0,0],
     [1,1,1,0]]
    We erase the underlined brick at (1,0), resulting in the grid:
    [[1,0,0,0],
     [0,1,1,0]]
    The two underlined bricks are no longer stable as they are no longer connected to the top nor adjacent to another stable brick, so they will fall. The resulting grid is:
    [[1,0,0,0],
     [0,0,0,0]]
    Hence the result is [2].

  Example 2:
    Input: grid = [[1,0,0,0],[1,1,0,0]], hits = [[1,1],[1,0]]
    Output: [0,0]
    Explanation: Starting with the grid:
    [[1,0,0,0],
     [1,1,0,0]]
    We erase the underlined brick at (1,1), resulting in the grid:
    [[1,0,0,0],
     [1,0,0,0]]
    All remaining bricks are still stable, so no bricks fall. The grid remains the same:
    [[1,0,0,0],
     [1,0,0,0]]
    Next, we erase the underlined brick at (1,0), resulting in the grid:
    [[1,0,0,0],
     [0,0,0,0]]
    Once again, all remaining bricks are still stable, so no bricks fall.
    Hence the result is [0,0].

  Constraints:
    m == grid.length
    n == grid[i].length
    1 <= m, n <= 200
    grid[i][j] is 0 or 1.
    1 <= hits.length <= 4 * 104
    hits[i].length == 2
    0 <= xi <= m - 1
    0 <= yi <= n - 1
    All (xi, yi) are unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/bricks-falling-when-hit
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find + BFS
func hitBricks(grid [][]int, hits [][]int) []int {
	m, n, h := len(grid), len(grid[0]), len(hits)
	out := make([]int, h)

	// hit all bricks in hits array
	for k, hit := range hits {
		if grid[hit[0]][hit[1]] == 0 {
			continue
		}
		out[k] = -1
		grid[hit[0]][hit[1]] = 0
	}

	count := 0
	queue := [][]int{}

	// put all top brick in queue
	for j := 0; j < n; j++ {
		if grid[0][j] == 0 {
			continue
		}
		count++
		queue = append(queue, []int{0, j})
	}

	uf := newUnionFind(m * n)

	// union all stable brick
	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		idx := i*n + j
		queue = queue[1:]
		for _, move := range moves {
			ni, nj := i+move[0], j+move[1]
			newIdx := ni*n + nj
			if ni < 1 || m-1 < ni || nj < 0 || n-1 < nj || grid[ni][nj] == 0 || uf.find(newIdx)/n == 0 {
				continue
			}
			uf.union(newIdx, idx)
			count++
			queue = append(queue, []int{ni, nj})
		}
	}

	// restore broken bricks in reverse order
	for k := h - 1; k >= 0; k-- {
		if out[k] != -1 {
			continue
		}

		i, j := hits[k][0], hits[k][1]
		idx := i*n + j
		grid[i][j] = 1
		cur := count + 1

		// if restore's brick near to stable brick, union it
		for _, move := range moves {
			ni, nj := i+move[0], j+move[1]
			newIdx := ni*n + nj
			if ni < 0 || m-1 < ni || nj < 0 || n-1 < nj || grid[ni][nj] == 0 || uf.find(newIdx)/n > 0 {
				continue
			}
			uf.union(idx, newIdx)
			break
		}

		queue := [][]int{}
		if uf.find(idx)/n == 0 {
			queue = append(queue, []int{i, j})
		}

		// union all bricks that are not stable
		for len(queue) > 0 {
			i, j := queue[0][0], queue[0][1]
			idx := i*n + j
			queue = queue[1:]
			for _, move := range moves {
				ni, nj := i+move[0], j+move[1]
				newIdx := ni*n + nj
				if ni < 1 || m-1 < ni || nj < 0 || n-1 < nj || grid[ni][nj] == 0 || uf.find(newIdx)/n == 0 {
					continue
				}
				uf.union(newIdx, idx)
				cur++
				queue = append(queue, []int{ni, nj})
			}
		}

		out[k] = cur - count - 1
		count = cur
	}

	return out
}

var moves = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type unionFind struct {
	ancestor []int
	isEnd    []bool
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := range ancestor {
		ancestor[i] = i
	}
	isEnd := make([]bool, n)
	return unionFind{ancestor: ancestor, isEnd: isEnd}
}

func (uf unionFind) find(x int) int {
	if uf.isEnd[uf.ancestor[x]] {
		return uf.ancestor[x]
	}
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
		uf.isEnd[x] = false
		uf.isEnd[uf.ancestor[x]] = true
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(x, y int) {
	ancestorX := uf.find(x)
	ancestorY := uf.find(y)
	uf.ancestor[ancestorX] = ancestorY
	uf.isEnd[ancestorX] = false
	uf.isEnd[ancestorY] = true
}
