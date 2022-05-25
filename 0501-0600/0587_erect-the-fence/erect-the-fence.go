package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Monotone Chain
// Sorting

func outerTrees(trees [][]int) [][]int {
	sort.Slice(trees, func(i, j int) bool {
		a, b := trees[i], trees[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
	})
	stk := [][]int{}
	top := func() int { return len(stk) - 1 }
	cross := func(o, a, b []int) bool { return (o[1]-a[1])*(o[0]-b[0]) > (o[1]-b[1])*(o[0]-a[0]) }
	for _, t := range trees {
		for ; len(stk) > 1 && cross(stk[top()-1], stk[top()], t); stk = stk[:top()] {
		}
		stk = append(stk, t)
	}
	for k, i := len(stk), len(trees)-2; i >= 0; i-- {
		for ; len(stk) > k && cross(stk[top()-1], stk[top()], trees[i]); stk = stk[:top()] {
		}
		stk = append(stk, trees[i])
	}
	vis := [101][101]bool{}
	i := 0
	for _, t := range stk {
		x, y := t[0], t[1]
		if vis[x][y] {
			continue
		}
		vis[x][y] = true
		stk[i] = t
		i++
	}
	return stk[:i]
}
