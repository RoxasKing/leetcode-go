package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Monotone Chain
// Sorting

func outerTrees(trees [][]int) [][]int {
	n := len(trees)
	out := make([][]int, 0, n<<1)
	top := func() int { return len(out) - 1 }
	sort.Slice(trees, func(i, j int) bool {
		return trees[i][0] < trees[j][0] || trees[i][0] == trees[j][0] && trees[i][1] < trees[j][1]
	})
	for i := 0; i < n; i++ {
		for ; len(out) > 1 && cross(out[top()-1], out[top()], trees[i]); out = out[:top()] {
		}
		out = append(out, trees[i])
	}
	for i, k := n-2, len(out); i >= 0; i-- {
		for ; len(out) > k && cross(out[top()-1], out[top()], trees[i]); out = out[:top()] {
		}
		out = append(out, trees[i])
	}
	vis := [101][101]bool{}
	i := 0
	for j := 0; j < len(out); j++ {
		x, y := out[j][0], out[j][1]
		if vis[x][y] {
			continue
		}
		vis[x][y] = true
		out[i] = out[j]
		i++
	}
	return out[:i]
}

func cross(o, a, b []int) bool { return (a[1]-o[1])*(b[0]-o[0]) > (b[1]-o[1])*(a[0]-o[0]) }
