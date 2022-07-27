package main

// Difficulty:
// Medium

// Tags:
// Topological Sort

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	n := len(nums)
	g := make([][]int, n+1)
	d := make([]int, n+1)
	for _, seq := range sequences {
		u := seq[0]
		for _, v := range seq[1:] {
			g[u] = append(g[u], v)
			d[v]++
			u = v
		}
	}
	q := []int{}
	for u := 1; u <= n; u++ {
		if d[u] == 0 {
			if q = append(q, u); len(q) > 1 {
				return false
			}
		}
	}
	if len(q) == 0 {
		return false
	}
	for ; len(q) > 0; q = q[1:] {
		u, c := q[0], 0
		if nums[0] != u {
			return false
		}
		nums = nums[1:]
		for _, v := range g[u] {
			if d[v]--; d[v] == 0 {
				if c++; c > 1 {
					return false
				}
				q = append(q, v)
			}
		}
	}
	return true
}
