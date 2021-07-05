package main

// Tags:
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
